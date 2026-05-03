package owner

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) GetByIdentification(ctx context.Context, identification string) (models.OwnerBO, error) {
	return s.repo.GetByIdentification(ctx, identification)
}
