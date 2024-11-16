package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/database"
	"user-service/domain"
	"user-service/domain/mappings"
	"user-service/tools/usercontext"
)

type User interface {
	GetUser(ctx usercontext.UserContext, id string) (domain.User, error)
	GetUsers(ctx usercontext.UserContext, limit int64, cursor string) ([]domain.User, error)
	SaveUser(ctx usercontext.UserContext, user domain.User) error
	DeleteUser(ctx usercontext.UserContext, id string) error
	UpdateUser(ctx usercontext.UserContext, user domain.User) error
}

type Impl struct {
	userRepository database.UserRepository
}

func NewUserService(repository database.UserRepository) User {
	return &Impl{
		userRepository: repository,
	}
}

func (s *Impl) GetUser(ctx usercontext.UserContext, id string) (domain.User, error) {
	dbUser, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, domain.NoDocuments
		}

		return domain.User{}, err
	}

	return mappings.ToDomain(dbUser), err
}

func (s *Impl) GetUsers(ctx usercontext.UserContext, limit int64, cursor string) ([]domain.User, error) {
	dbUsers, err := s.userRepository.GetUsers(ctx, limit, cursor)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return []domain.User{}, domain.NoDocuments
		}

		return []domain.User{}, err
	}

	return mappings.ToDomainSlice(dbUsers), err
}

func (s *Impl) SaveUser(ctx usercontext.UserContext, user domain.User) error {
	return s.userRepository.SaveUsers(ctx, []database.User{mappings.ToDatabase(user)})
}

func (s *Impl) DeleteUser(ctx usercontext.UserContext, id string) error {
	return s.userRepository.DeleteUsers(ctx, []string{id})
}

func (s *Impl) UpdateUser(ctx usercontext.UserContext, user domain.User) error {
	return s.userRepository.UpdateUsers(ctx, []database.User{mappings.ToDatabase(user)})
}
