package assembly

import (
	"asamblea/internal/domain"
	"asamblea/internal/models"
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
)

func generateCode() string {
	b := make([]byte, 4)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (r *repository) GenerateCodes(ctx context.Context, assemblyID int, units []models.UnitBO) error {
	for _, unit := range units {
		// verificar si ya existe código para esta unidad en esta asamblea
		var existing models.AssemblyCodeEntity
		err := r.db.Where("assembly_id = ? AND unit_id = ?", assemblyID, unit.ID).First(&existing).Error
		if err == nil {
			continue // ya tiene código
		}

		entity := models.AssemblyCodeEntity{
			AssemblyID: assemblyID,
			UnitID:     unit.ID,
			Code:       generateCode(),
			Used:       false,
		}
		if err = r.db.Create(&entity).Error; err != nil {
			derr := domain.SavingErr("código")
			slog.Error(derr.Error(), "err", err)
			return derr
		}
	}
	return nil
}

func (r *repository) GetCode(ctx context.Context, code string) (models.AssemblyCodeBO, error) {
	var entity models.AssemblyCodeEntity
	err := r.db.Preload("Unit").Where("code = ?", code).First(&entity).Error
	if err != nil {
		return models.AssemblyCodeBO{}, domain.NotFounErr("código")
	}
	return models.AssemblyCodeBO{
		ID:         entity.ID,
		AssemblyID: entity.AssemblyID,
		UnitID:     entity.UnitID,
		UnitNumber: entity.Unit.Number,
		Code:       entity.Code,
		Used:       entity.Used,
	}, nil
}

func (r *repository) MarkCodeUsed(ctx context.Context, code string) error {
	err := r.db.Model(&models.AssemblyCodeEntity{}).
		Where("code = ?", code).
		Update("used", true).Error
	if err != nil {
		return domain.SavingErr("código")
	}
	return nil
}

func (r *repository) GetCodes(ctx context.Context, assemblyID int) ([]models.AssemblyCodeBO, error) {
	var entities []models.AssemblyCodeEntity
	err := r.db.Preload("Unit").Where("assembly_id = ?", assemblyID).Find(&entities).Error
	if err != nil {
		return nil, domain.NotFounErr("códigos")
	}
	result := make([]models.AssemblyCodeBO, len(entities))
	for i, e := range entities {
		result[i] = models.AssemblyCodeBO{
			ID:         e.ID,
			AssemblyID: e.AssemblyID,
			UnitID:     e.UnitID,
			UnitNumber: e.Unit.Number,
			Code:       e.Code,
			Used:       e.Used,
		}
	}
	return result, nil
}
