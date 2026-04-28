package apartment

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) Get(ctx context.Context, id int) (models.ApartmentBO, error) {

	apartment, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.ApartmentBO{}, err
	}
	return *apartment, nil
}
