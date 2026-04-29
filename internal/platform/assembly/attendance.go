package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

func (r *repository) RegisterAttendance(ctx context.Context, a models.AssemblyUnitBO) error {
	entity := mapAttendanceToEntity(a)
	if err := r.db.Save(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.DuplicatedErr("asistencia")
		}
		derr := domain.SavingErr("asistencia")
		slog.Error(derr.Error(), "err", err)
		return derr
	}
	return nil
}

func (r *repository) GetAttendance(ctx context.Context, assemblyID int) ([]models.AssemblyUnitBO, error) {
	var entities []models.AssemblyUnitEntity
	err := r.db.Preload("Unit").Where("assembly_id = ?", assemblyID).Find(&entities).Error
	if err != nil {
		derr := domain.NotFounErr("asistencia")
		slog.Error(derr.Error(), "err", err)
		return nil, derr
	}
	result := make([]models.AssemblyUnitBO, len(entities))
	for i, e := range entities {
		result[i] = mapEntityToAttendance(e)
	}
	return result, nil
}
