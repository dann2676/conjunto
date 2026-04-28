package owner

import (
	"context"
)

func (s *service) Delete(ctx context.Context, id int) error {

	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
