package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *repository) Save(ctx context.Context, owner models.OwnerBO) error {
	entity := mapBOToEntity(owner)
	if err := r.db.Save(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.DuplicatedErr("dueño")
		}
		return domain.SavingErr("dueño")
	}
	return nil
}
