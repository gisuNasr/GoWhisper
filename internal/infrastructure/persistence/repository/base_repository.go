package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/gisuNasr/GoWhisper/internal/infrastructure/persistence/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepository[TEntity any] struct {
	DB *gorm.DB
}

func NewBaseRepository[TEntity any]() *BaseRepository[TEntity] {
	return &BaseRepository[TEntity]{
		DB: database.GetDB(),
	}
}

func (r BaseRepository[TEntity]) Create(ctx context.Context, entity TEntity) (TEntity, error) {
	err := r.DB.WithContext(ctx).Create(&entity).Error
	return entity, err
}

func (r BaseRepository[TEntity]) Update(ctx context.Context, id uuid.UUID, updateData map[string]interface{}) (TEntity, error) {
	updateData["updated_at"] = time.Now().UTC()

	model := new(TEntity)

	err := r.DB.WithContext(ctx).
		Model(model).
		Where("id = ?", id).
		Updates(updateData).
		Error

	if err != nil {
		return *model, err
	}

	err = r.DB.WithContext(ctx).Where("id = ?", id).First(model).Error
	return *model, err
}

func (r BaseRepository[TEntity]) Delete(ctx context.Context, id uuid.UUID) error {
	model := new(TEntity)

	result := r.DB.WithContext(ctx).Where("id = ?", id).Delete(model)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}

	return nil
}

func (r BaseRepository[TEntity]) GetById(ctx context.Context, id uuid.UUID) (TEntity, error) {
	model := new(TEntity)

	err := r.DB.WithContext(ctx).
		Where("id = ?", id).
		First(model).
		Error

	return *model, err
}

func (r BaseRepository[TEntity]) GetByFilter(
	ctx context.Context,
	limit int,
	offset int,
	filters map[string]interface{},
) (int64, []TEntity, error) {

	model := new(TEntity)
	var items []TEntity
	var totalRows int64

	query := r.DB.WithContext(ctx).Model(model)

	if len(filters) > 0 {
		query = query.Where(filters)
	}

	if err := query.Count(&totalRows).Error; err != nil {
		return 0, nil, err
	}

	err := query.
		Offset(offset).
		Limit(limit).
		Find(&items).
		Error

	if err != nil {
		return 0, nil, err
	}

	return totalRows, items, nil
}
