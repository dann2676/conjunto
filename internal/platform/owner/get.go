package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) Get(ctx context.Context, id int) (*models.OwnerBO, error) {
	var owner models.OwnerEntity
	err := r.db.Find(&owner, "id = ?", id).Error
	if err != nil {
		derr := domain.SavingErr("dueño")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}
	return mapEntityToBO(owner), nil
}
