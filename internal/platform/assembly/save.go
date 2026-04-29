package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) Save(ctx context.Context, assembly models.AssemblyBO) error {
	entity := mapAssemblyToEntity(assembly)
	if err := r.db.Save(&entity).Error; err != nil {
		derr := domain.SavingErr("asamblea")
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}
