package mappings

import (
	"github.com/samber/lo"
	"user-service/database"
	"user-service/domain"
)

func ToDomain(user database.User) domain.User {
	return domain.User{
		Id:       user.Id,
		Surname:  user.Surname,
		Name:     user.Name,
		Lastname: user.Lastname,
		Role:     domain.Role(user.Role),
	}
}

func ToDatabase(user domain.User) database.User {
	return database.User{
		Id:       user.Id,
		Surname:  user.Surname,
		Name:     user.Name,
		Lastname: user.Lastname,
		Role:     uint(user.Role),
	}
}

func ToDomainSlice(users []database.User) []domain.User {
	return lo.Map(users, func(item database.User, _ int) domain.User {
		return ToDomain(item)
	})
}

func ToDatabaseSlice(users []domain.User) []database.User {
	return lo.Map(users, func(item domain.User, _ int) database.User {
		return ToDatabase(item)
	})
}
