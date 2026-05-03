package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *repository) GetBySlug(ctx context.Context, slug string) (models.AssemblyBO, error) {
	var entity models.AssemblyEntity
	err := r.db.First(&entity, "slug = ?", slug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.AssemblyBO{}, domain.NotFounErr("asamblea")
		}
		return models.AssemblyBO{}, domain.SavingErr("asamblea")
	}
	return mapEntityToAssembly(entity), nil
}
