package assembly

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) Update(ctx context.Context, assembly models.AssemblyBO) error {
	existing, err := s.repo.Get(ctx, assembly.ID)
	if err != nil {
		return err
	}
	assembly.Status = existing.Status
	assembly.Status = existing.Slug
	return s.repo.Save(ctx, assembly)
}
