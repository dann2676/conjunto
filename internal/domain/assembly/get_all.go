package assembly

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) GetAll(ctx context.Context, i bool) ([]models.AssemblyBO, error) {
	return s.repo.GetAll(ctx, i)
}
