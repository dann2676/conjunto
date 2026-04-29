package owner

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) Update(ctx context.Context, owner models.OwnerBO) error {
	existing, err := s.repo.Get(ctx, owner.ID)
	if err != nil {
		return err
	}
	owner.StartDate = existing.StartDate
	return s.repo.Save(ctx, owner)
}
