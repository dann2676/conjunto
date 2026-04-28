package unit

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) Get(ctx context.Context, id int) (models.UnitBO, error) {
	var unit models.UnitEntity
	err := r.db.Find(&unit, "id = ?", id).Error
	if err != nil {
		derr := domain.NotFounErr("unito")
		slog.Error(derr.Error(), "err", err)
		return models.UnitBO{}, derr
	}
	return mapEntityToBO(unit), nil
}
