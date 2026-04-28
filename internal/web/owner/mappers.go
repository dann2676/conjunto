package owner

import (
	"asamblea/internal/models"
)

func mapRequestToBO(owner models.OwnerRequest) models.OwnerBO {
	return models.OwnerBO{
		Name:        owner.Name,
		Email:       owner.Email,
		Phone:       owner.Phone,
		ApartmentID: owner.ApartmentID,
	}
}
func mapBOToDTO(owner models.OwnerBO) models.OwnerDTO {
	return models.OwnerDTO{
		ID:          owner.ID,
		Apartment:   owner.Apartment,
		ApartmentID: owner.ApartmentID,
		Name:        owner.Name,
		Email:       owner.Email,
		Phone:       owner.Phone,
	}
}

func mapBOsToDTOs(owners []models.OwnerBO) []models.OwnerDTO {
	response := make([]models.OwnerDTO, len(owners))
	for i, owner := range owners {
		response[i] = mapBOToDTO(owner)
	}
	return response
}
