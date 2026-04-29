package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

func (r *repository) Get(ctx context.Context, id int) (models.AssemblyBO, error) {
	var entity models.AssemblyEntity
	err := r.db.First(&entity, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.AssemblyBO{}, domain.NotFounErr("asamblea")
		}
		slog.Error("error getting assembly", "err", err)
		return models.AssemblyBO{}, domain.SavingErr("asamblea")
	}
	return mapEntityToAssembly(entity), nil
}
