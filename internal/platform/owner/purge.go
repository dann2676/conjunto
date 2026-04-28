package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) Purge(ctx context.Context, id int) error {
	err := r.db.Unscoped().Delete(&models.OwnerEntity{}, id).
		Error
	if err != nil {
		derr := domain.DeletingErr("dueño")
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}
