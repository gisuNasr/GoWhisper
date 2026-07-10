package service

import (
	"context"

	"github.com/gisuNasr/GoWhisper/internal/application/dto"
	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) Create(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	if req.Username == "" {
		return nil, domain.ErrInvalidInput
	}

	user := &domain.User{
		UserName:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	user.ID = uuid.New()

	if err := s.UserRepository.Create(ctx, user); err != nil {
		return nil, err
	}
	return dto.ToUserResponse(user), nil
}

func (s *UserService) GetByID(ctx context.Context, id uuid.UUID) (*dto.UserResponse, error) {
	user, err := s.UserRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToUserResponse(user), nil
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*dto.UserResponse, error) {
	user, err := s.UserRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return dto.ToUserResponse(user), nil
}
