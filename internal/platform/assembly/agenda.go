package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"log/slog"
)

func (r *repository) GetAgendaItem(ctx context.Context, itemID int) (models.AgendaItemBO, error) {
	var entity models.AgendaItemEntity
	err := r.db.First(&entity, "id = ?", itemID).Error
	if err != nil {
		return models.AgendaItemBO{}, domain.NotFounErr("agenda")
	}
	return mapEntityToAgendaItem(entity), nil
}
func (r *repository) GetAgendaItems(ctx context.Context, assemblyID int) ([]models.AgendaItemBO, error) {
	var entities []models.AgendaItemEntity
	err := r.db.Where("assembly_id = ?", assemblyID).Order(`"order"`).Find(&entities).Error
	if err != nil {
		derr := domain.NotFounErr("agenda")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}
	result := make([]models.AgendaItemBO, len(entities))
	for i, e := range entities {
		result[i] = mapEntityToAgendaItem(e)
	}
	return result, nil
}

func (r *repository) SaveAgendaItem(ctx context.Context, item models.AgendaItemBO) error {
	entity := mapAgendaItemToEntity(item)
	if err := r.db.Save(&entity).Error; err != nil {
		derr := domain.SavingErr("agenda")
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}

func (r *repository) DeleteAgendaItem(ctx context.Context, itemID int) error {
	err := r.db.Delete(&models.AgendaItemEntity{}, itemID).Error
	if err != nil {
		derr := domain.DeletingErr("agenda")
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}
