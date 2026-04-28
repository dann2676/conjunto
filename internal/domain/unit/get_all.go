package unit

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) GetAll(ctx context.Context, includeInactive bool) ([]models.UnitBO, error) {
	return s.repo.GetAll(ctx, includeInactive)
}
