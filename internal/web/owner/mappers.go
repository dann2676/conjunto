package owner

import (
	"asamblea/internal/models"
)

func mapRequestToBO(owner models.OwnerRequest) models.OwnerBO {
	return models.OwnerBO{
		Name:           owner.Name,
		Email:          owner.Email,
		Phone:          owner.Phone,
		UnitID:         owner.UnitID,
		Identification: owner.Identification,
	}
}
func mapBOToDTO(owner models.OwnerBO) models.OwnerDTO {
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

func mapBOsToDTOs(owners []models.OwnerBO) []models.OwnerDTO {
	response := make([]models.OwnerDTO, len(owners))
	for i, owner := range owners {
		response[i] = mapBOToDTO(owner)
	}
	return response
}
