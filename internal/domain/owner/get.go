package owner

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"
)

func (s *service) Get(ctx context.Context, id int) (models.OwnerBO, error) {

	owner, err := s.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, domain.NotFounErr("")) {
			return models.OwnerBO{}, nil
		}
		return models.OwnerBO{}, err
	}
	return *owner, nil
}
