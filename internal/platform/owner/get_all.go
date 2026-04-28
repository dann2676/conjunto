package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) GetAll(ctx context.Context) ([]models.OwnerBO, error) {
	var entidades []models.OwnerEntity

	err := r.db.Preload("Apartment").Where("active", 1).Order("name").Find(&entidades).Error
	if err != nil {
		derr := domain.NotFounErr("dueños")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}

	return mapEntitiesToBOs(entidades), nil
}
