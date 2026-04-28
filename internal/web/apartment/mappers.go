package apartment

import (
	"asamblea/internal/models"
)

func mapRequestToBO(aparment models.ApartmentRequest) models.ApartmentBO {
	return models.ApartmentBO{
		Number:     aparment.Number,
		Coeficient: aparment.Coeficient,
		Meters:     aparment.Meters,
	}
}
func mapBOToDTO(apartment models.ApartmentBO) models.ApartmentDTO {
	return models.ApartmentDTO{
		ID:         apartment.ID,
		Number:     apartment.Number,
		Coeficient: apartment.Coeficient,
		Meters:     apartment.Meters,
	}
}

func mapBOsToDTOs(apartments []models.ApartmentBO) []models.ApartmentDTO {
	response := make([]models.ApartmentDTO, len(apartments))
	for i, apartment := range apartments {
		response[i] = mapBOToDTO(apartment)
	}
	return response
}
