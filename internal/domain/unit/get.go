package unit

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) Get(ctx context.Context, id int) (models.UnitBO, error) {

	unit, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.UnitBO{}, err
	}
	return unit, nil
}
