package service

import (
	"context"
	"go_test/src/domain/user/entity"
	"go_test/src/domain/user/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{userRepository: repository}
}

func (s *UserService) CreateUser(ctx context.Context, name string, email string) (*entity.User, error) {
	newUser := entity.User{
		Name:  name,
		Email: email,
	}

	return s.userRepository.CreateUser(&newUser)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, name *string, email *string) (*entity.User, error) {
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	if name != nil {
		user.Name = *name
	}
	if email != nil {
		user.Email = *email
	}

	return s.userRepository.UpdateUser(user)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) (*entity.User, error) {
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	if err := s.userRepository.DeleteUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return s.userRepository.GetUser(id)
}

func (s *UserService) GetUsers(ctx context.Context) (*[]entity.User, error) {
	return s.userRepository.GetUsers()
}
