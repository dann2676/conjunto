package owner

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) GetActiveByUnit(ctx context.Context, unitID int) (models.OwnerBO, error) {
	return s.repo.GetActiveByUnit(ctx, unitID)
}
