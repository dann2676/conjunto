package assembly

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) GetBySlug(ctx context.Context, slug string) (models.AssemblyBO, error) {
	return s.repo.GetBySlug(ctx, slug)
}
