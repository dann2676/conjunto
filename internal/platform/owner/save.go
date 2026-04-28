package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (r *repository) Save(ctx context.Context, owner models.OwnerBO) error {
	entity := mapBOToEntity(owner)

	if entity.ID != 0 {
		var existing models.OwnerEntity
		err := r.db.First(&existing, "id = ?", entity.ID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return domain.NotFounErr("dueño")
			}
			return domain.SavingErr("dueño")
		}
		entity.StartDate = existing.StartDate
	} else {
		entity.StartDate = time.Now()
	}

	if err := r.db.Save(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.DuplicatedErr("dueño")
		}
		return domain.SavingErr("dueño")
	}
	return nil
}
