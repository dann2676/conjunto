package assembly

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) Get(ctx context.Context, id int) (models.AssemblyBO, error) {
	return s.repo.Get(ctx, id)
}
