package kafka

import (
	"encoding/json"
	"fmt"
	"time"
	"user-service/database"
	"user-service/tools/usercontext"
)

var (
	location, _ = time.LoadLocation("Europe/Moscow")
)

type MessageHandler interface {
	HandleMessage(_ usercontext.UserContext, _ []byte) error
}

type Impl struct {
	userRepository database.UserRepository
	kafkaProducer  *Producer
}

func NewMessageHandler(userRepository database.UserRepository, producer *Producer) MessageHandler {
	return &Impl{
		userRepository: userRepository,
		kafkaProducer:  producer,
	}
}

func (h *Impl) HandleMessage(ctx usercontext.UserContext, message []byte) error {
	var approveItems []database.ApproveMessage
	if err := json.Unmarshal(message, &approveItems); err != nil {
		return fmt.Errorf("не удалось обработать сообщение: %w", err)
	}

	approvedItems := make([]database.ApprovedItem, 0, len(approveItems))

	for _, item := range approveItems {
		ctx.Log().Info(fmt.Sprintf("Продукт %s ожидает подтвержения от %s", item.ProductId, item.UserId))

		approver, err := h.userRepository.GetUser(ctx, item.UserId)
		if err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось получить согласующего из базы: %v", err))
			continue
		}

		approver.RegisteredObjects++

		if err = h.userRepository.UpdateUsers(ctx, []database.User{approver}); err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось обновить информацию о согласующем: %v", err))
			continue
		}

		approvedItems = append(approvedItems, database.ApprovedItem{
			ProductId:   item.ProductId,
			ApproveTime: time.Now(),
		})

		ctx.Log().Info(fmt.Sprintf("Продукт %s подтвержден %s", item.ProductId, item.UserId))
	}

	if err := h.kafkaProducer.Produce(approvedItems); err != nil {
		return fmt.Errorf("не удалось отпраить сообщение в кафку: %w", err)
	}

	return nil
}
