package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
	"time"
)

func (r *repository) Delete(ctx context.Context, id int) error {
	err := r.db.Model(&models.OwnerEntity{}).Where("id = ?", id).
		Update("active", 0).
		Update("end_date", time.Now()).
		Update("deleted_at", time.Now()).
		Error
	if err != nil {
		derr := domain.DeletingErr("dueño")
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}
