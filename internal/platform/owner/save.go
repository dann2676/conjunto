package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
	"time"
)

func (r *repository) Save(ctx context.Context, owner models.OwnerBO) error {
	entity := mapBOToEntity(owner)
	entity.Active = true
	if entity.ID != 0 {
		var existing models.OwnerEntity
		err := r.db.Find(&existing, "id = ?", entity.ID).Error
		if err != nil {
			derr := domain.SavingErr("dueño")
			slog.Error(derr.Error(), "err", err)
			return derr
		}
		slog.Info("i, entonce", "existing", existing, "id", entity.ID)
		entity.StartDate = existing.StartDate
	} else {
		entity.StartDate = time.Now()
	}

	slog.Debug("saving", "data", owner, "entity", entity)
	err := r.db.Save(&entity).Error
	if err != nil {
		derr := domain.SavingErr("dueño")
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}
