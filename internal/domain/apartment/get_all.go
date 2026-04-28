package apartment

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) GetAll(ctx context.Context) ([]models.ApartmentBO, error) {

	return s.repo.GetAll(ctx)
}
