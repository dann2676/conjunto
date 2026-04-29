package owner

import "asamblea/internal/models"

func mapBOToEntity(owner models.OwnerBO) models.OwnerEntity {
	o := models.OwnerEntity{
		ID:             owner.ID,
		Name:           owner.Name,
		Email:          owner.Email,
		Phone:          owner.Phone,
		UnitID:         owner.UnitID,
		Identification: owner.Identification,
		StartDate:      owner.StartDate,
	}
	return o

}
func mapEntityToBO(owner models.OwnerEntity) models.OwnerBO {
	return models.OwnerBO{
		ID:             owner.ID,
		Identification: owner.Identification,
		Name:           owner.Name,
		Email:          owner.Email,
		Phone:          owner.Phone,
		UnitID:         owner.UnitID,
		Unit:           owner.Unit.Number,
		Active:         !owner.DeletedAt.Valid,
		StartDate:      owner.StartDate,
	}
}

func mapEntitiesToBOs(owners []models.OwnerEntity) []models.OwnerBO {
	result := make([]models.OwnerBO, len(owners))
	for i, owner := range owners {
		result[i] = mapEntityToBO(owner)
	}
	return result
}
