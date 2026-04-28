package apartment

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
	"strings"
)

func (r *repository) Delete(ctx context.Context, id int) error {
	err := r.db.Delete(&models.ApartmentEntity{}, id).Error
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
