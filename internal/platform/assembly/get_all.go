package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) GetAll(ctx context.Context, _ bool) ([]models.AssemblyBO, error) {
	var entities []models.AssemblyEntity
	err := r.db.Order("date desc").Find(&entities).Error
	if err != nil {
		derr := domain.NotFounErr("asambleas")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}
	return mapEntitiesToAssemblies(entities), nil
}
