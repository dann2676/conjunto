package apartment

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) GetAll(ctx context.Context) ([]models.ApartmentBO, error) {
	var entidades []models.ApartmentEntity

	err := r.db.Order("number").Find(&entidades).Error
	if err != nil {
		derr := domain.NotFounErr("apartmentos")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}

	return mapEntitiesToBOs(entidades), err
}
