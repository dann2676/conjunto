package owner

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) GetAll(ctx context.Context, includeInactive bool) ([]models.OwnerBO, error) {
	return s.repo.GetAll(ctx, includeInactive)
}
