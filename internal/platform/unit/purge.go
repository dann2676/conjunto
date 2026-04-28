package unit

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
	"strings"
)

func (r *repository) Purge(ctx context.Context, id int) error {
	err := r.db.Unscoped().Delete(&models.UnitEntity{}, id).Error
	if err != nil {
		derr := domain.DeletingErr("apartamento")
		if strings.Contains(err.Error(), "FOREIGN KEY constraint failed") {
			derr = domain.AsociationErr("apartamento")
		}

		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}
