package repository

import (
	"context"
	"encoding/json"

	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type deviceRepository struct {
	*BaseRepository[domain.Device]
}

func NewDeviceRepository() domain.DeviceRepository {
	return &deviceRepository{
		NewBaseRepository[domain.Device](),
	}
}

func (r *deviceRepository) Create(ctx context.Context, device *domain.Device) error {
	createdModel, err := r.BaseRepository.Create(ctx, *device)
	if err != nil {
		return err
	}
	*device = createdModel
	return nil
}

func (r *deviceRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Device, error) {
	count, items, err := r.BaseRepository.GetByFilter(ctx, 15, 0, map[string]interface{}{"user_id": userID})
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}
	devices := make([]*domain.Device, len(items))
	for i := range items {
		devices[i] = &items[i]

	}
	return devices, nil
}

func (r *deviceRepository) Delete(ctx context.Context, deviceID uuid.UUID) error {
	err := r.BaseRepository.Delete(ctx, deviceID)
	return err
}

func (r *deviceRepository) UpdateSignedPreKey(ctx context.Context, deviceID uuid.UUID, newPubKey string) error {
	_, err := r.BaseRepository.Update(ctx, deviceID, map[string]interface{}{"signed_pre_key_pub": newPubKey})
	return err
}

func (r *deviceRepository) ClaimOneTimePreKey(ctx context.Context, deviceID uuid.UUID) (string, error) {
	var claimedKey string

	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var rawKeys []byte

		err := tx.Model(&domain.Device{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Select("one_time_pre_keys").
			Where("id = ?", deviceID).
			Scan(&rawKeys).Error

		if err != nil {
			return err
		}

		if len(rawKeys) == 0 {
			return domain.ErrNoOneTimePreKeys
		}

		var keys []string
		if err := json.Unmarshal(rawKeys, &keys); err != nil {
			return err
		}
		if len(keys) == 0 {
			return domain.ErrNoOneTimePreKeys
		}

		claimedKey = keys[0]
		remaining := keys[1:]

		updatedJSON, err := json.Marshal(remaining)
		if err != nil {
			return err
		}

		err = tx.Model(&domain.Device{}).
			Where("id = ?", deviceID).
			Update("one_time_pre_keys", updatedJSON).Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return claimedKey, nil
}

func (r *deviceRepository) AddOneTimePreKeys(ctx context.Context, deviceID uuid.UUID, newKeys []string) error {

	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var rawKeys []byte

		err := tx.Model(&domain.Device{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Select("one_time_pre_keys").
			Where("id = ?", deviceID).
			Scan(&rawKeys).Error

		if err != nil {
			return err
		}

		var existing []string

		if len(rawKeys) > 0 {
			if err := json.Unmarshal(rawKeys, &existing); err != nil {
				return err
			}
		}

		merged := append(existing, newKeys...)

		updatedJSON, err := json.Marshal(merged)
		if err != nil {
			return err
		}

		err = tx.Model(&domain.Device{}).
			Where("id = ?", deviceID).
			Update("one_time_pre_keys", updatedJSON).Error

		return err
	})
}
