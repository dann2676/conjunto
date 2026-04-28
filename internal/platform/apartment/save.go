package apartment

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

func (r *repository) Save(ctx context.Context, aparment models.ApartmentBO) error {
	entity := mapBOToEntity(aparment)
	slog.Debug("saving", "data", aparment, "entity", entity)
	err := r.db.Save(&entity).Error
	if err != nil {
		derr := domain.SavingErr("apartmento")
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			derr = domain.DuplicatedErr("apartamento")
		}
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}
