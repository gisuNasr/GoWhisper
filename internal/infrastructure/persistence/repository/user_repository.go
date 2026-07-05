package repository

import (
	"context"

	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
)

type userRepository struct {
	*BaseRepository[domain.User]
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[domain.User](),
	}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	createdUser, err := r.BaseRepository.Create(ctx, *user)
	if err != nil {
		return err
	}
	*user = createdUser
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := r.BaseRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := r.DB.WithContext(ctx).
		Where("user_name = ?", username).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
