package unit

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

func (r *repository) Save(ctx context.Context, unit models.UnitBO) error {
	entity := mapBOToEntity(unit)
	slog.Debug("saving", "data", unit, "entity", entity)
	err := r.db.Save(&entity).Error
	if err != nil {
		derr := domain.SavingErr("unito")
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			derr = domain.DuplicatedErr("apartamento")
		}
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}
