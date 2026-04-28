package apartment

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) Create(ctx context.Context, apartment models.ApartmentBO) error {

	err := s.repo.Save(ctx, apartment)
	if err != nil {
		return err
	}
	return nil
}
