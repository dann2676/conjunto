package apartment

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) Get(ctx context.Context, id int) (*models.ApartmentBO, error) {
	var apartment models.ApartmentEntity
	err := r.db.Find(&apartment, "id = ?", id).Error
	if err != nil {
		derr := domain.NotFounErr("apartmento")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}
	return mapEntityToBO(apartment), nil
}
