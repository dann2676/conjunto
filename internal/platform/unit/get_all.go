package unit

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) GetAll(ctx context.Context, includeInactive bool) ([]models.UnitBO, error) {
	var entidades []models.UnitEntity

	q := r.db.Order("number")
	if includeInactive {
		q = q.Unscoped()
	}

	err := q.Find(&entidades).Error
	if err != nil {
		derr := domain.NotFounErr("unitos")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}

	return mapEntitiesToBOs(entidades), err
}
