package owner

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) Update(ctx context.Context, owner models.OwnerBO) error {

	err := s.repo.Save(ctx, owner)
	if err != nil {
		return err
	}
	return nil
}
