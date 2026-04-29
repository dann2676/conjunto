package assembly

import (
	"asamblea/internal/models"
	"context"
)

func (s *service) RegisterAttendance(ctx context.Context, a models.AssemblyUnitBO) error {
	return s.repo.RegisterAttendance(ctx, a)
}

func (s *service) GetAttendance(ctx context.Context, assemblyID int) ([]models.AssemblyUnitBO, error) {
	return s.repo.GetAttendance(ctx, assemblyID)
}

func (s *service) GetQuorum(ctx context.Context, assemblyID int) (float32, error) {
	attendance, err := s.repo.GetAttendance(ctx, assemblyID)
	if err != nil {
		return 0, err
	}
	var total float32
	for _, a := range attendance {
		total += a.Coeficient
	}
	return total, nil
}
