package unit

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) Create(ctx context.Context, unit models.UnitBO) error {

	err := s.repo.Save(ctx, unit)
	if err != nil {
		return err
	}
	return nil
}
