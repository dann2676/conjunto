package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) GetUnitsByOwner(ctx context.Context, ownerID int) ([]models.UnitBO, error) {
	var entities []models.OwnerEntity
	err := r.db.Preload("Unit").Where("id = ?", ownerID).Find(&entities).Error
	if err != nil {
		derr := domain.NotFounErr("unidades")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}
	units := make([]models.UnitBO, len(entities))
	for i, e := range entities {
		units[i] = models.UnitBO{
			ID:         e.Unit.ID,
			Number:     e.Unit.Number,
			Coeficient: e.Unit.Coeficient,
			Meters:     e.Unit.Meters,
			Type:       e.Unit.Type,
		}
	}
	return units, nil
}
