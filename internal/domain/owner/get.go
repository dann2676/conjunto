package owner

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) Get(ctx context.Context, id int) (models.OwnerBO, error) {

	owner, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.OwnerBO{}, err
	}
	return *owner, nil
}
