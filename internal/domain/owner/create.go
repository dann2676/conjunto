package owner

import (
	"asamblea/internal/models"
	"context"
	"time"
)

func (s *service) Create(ctx context.Context, owner models.OwnerBO) error {
	owner.StartDate = time.Now()
	return s.repo.Save(ctx, owner)
}
