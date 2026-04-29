package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
)

func (s *service) GetAgendaItems(ctx context.Context, assemblyID int) ([]models.AgendaItemBO, error) {
	return s.repo.GetAgendaItems(ctx, assemblyID)
}

func (s *service) CreateAgendaItem(ctx context.Context, item models.AgendaItemBO) error {
	item.Status = "pending"
	return s.repo.SaveAgendaItem(ctx, item)
}

func (s *service) UpdateAgendaItemStatus(ctx context.Context, itemID int, newStatus string) error {
	item, err := s.repo.GetAgendaItem(ctx, itemID)
	if err != nil {
		return err
	}

	valid := map[string]string{
		"pending": "open",
		"open":    "closed",
	}
	if valid[item.Status] != newStatus {
		return domain.DuplicatedErr("transición de estado inválida")
	}

	item.Status = newStatus
	return s.repo.SaveAgendaItem(ctx, item)
}

func (s *service) DeleteAgendaItem(ctx context.Context, itemID int) error {
	return s.repo.DeleteAgendaItem(ctx, itemID)
}
