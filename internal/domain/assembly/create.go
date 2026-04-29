package assembly

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) Create(ctx context.Context, assembly models.AssemblyBO) error {
	assembly.Status = "draft"
	return s.repo.Save(ctx, assembly)
}
