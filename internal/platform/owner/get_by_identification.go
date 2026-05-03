package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *repository) GetByIdentification(ctx context.Context, identification string) (models.OwnerBO, error) {
	var entity models.OwnerEntity
	err := r.db.Preload("Unit").Where("identification = ?", identification).First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.OwnerBO{}, domain.NotFounErr("dueño")
		}
		return models.OwnerBO{}, domain.SavingErr("dueño")
	}
	return mapEntityToBO(entity), nil
}
