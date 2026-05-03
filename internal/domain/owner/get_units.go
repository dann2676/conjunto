package owner

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) GetUnitsByOwner(ctx context.Context, ownerID int) ([]models.UnitBO, error) {
	return s.repo.GetUnitsByOwner(ctx, ownerID)
}
