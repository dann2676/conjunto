package unit

import (
	"context"
)

func (s *service) Purge(ctx context.Context, id int) error {

	err := s.repo.Purge(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
