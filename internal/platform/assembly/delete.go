package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) Delete(ctx context.Context, id int) error {
	err := r.db.Delete(&models.AssemblyEntity{}, id).Error
	if err != nil {
		derr := domain.DeletingErr("asamblea")
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}
