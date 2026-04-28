package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) GetAll(ctx context.Context, includeInactive bool) ([]models.OwnerBO, error) {
	var entidades []models.OwnerEntity

	q := r.db.Preload("Apartment").Order("name")
	if !includeInactive {
		q = q.Where("active = ?", 1)
	}

	err := q.Find(&entidades).Error
	if err != nil {
		derr := domain.NotFounErr("dueños")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}
	return mapEntitiesToBOs(entidades), nil
}
