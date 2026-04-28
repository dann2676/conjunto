package unit

import (
	"asamblea/internal/models"
)

func mapRequestToBO(unit models.UnitRequest) models.UnitBO {
	return models.UnitBO{
		Number:     unit.Number,
		Coeficient: unit.Coeficient,
		Meters:     unit.Meters,
		Type:       unit.Type,
	}
}
func mapBOToDTO(unit models.UnitBO) models.UnitDTO {
	return models.UnitDTO{
		ID:         unit.ID,
		Number:     unit.Number,
		Coeficient: unit.Coeficient,
		Meters:     unit.Meters,
		Type:       unit.Type,
	}
}

func mapBOsToDTOs(units []models.UnitBO) []models.UnitDTO {
	response := make([]models.UnitDTO, len(units))
	for i, unit := range units {
		response[i] = mapBOToDTO(unit)
	}
	return response
}

func mapOwnerBOToDTO(owner models.OwnerBO) models.OwnerDTO {
	return models.OwnerDTO{
		ID:             owner.ID,
		Unit:           owner.Unit,
		UnitID:         owner.UnitID,
		Identification: owner.Identification,
		Name:           owner.Name,
		Email:          owner.Email,
		Phone:          owner.Phone,
		Active:         owner.Active,
	}
}
