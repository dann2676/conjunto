package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

func (r *repository) Get(ctx context.Context, id int) (models.OwnerBO, error) {
	var owner models.OwnerEntity
	err := r.db.First(&owner, "id = ?", id).Error
	if err != nil {
		derr := domain.SavingErr("dueño")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			derr = domain.NotFounErr("dueño")
		}
		slog.Error(derr.Error(), "err", err)
		return models.OwnerBO{}, derr
	}
	return mapEntityToBO(owner), nil
}
