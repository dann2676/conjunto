package apartment

import "asamblea/internal/models"

func mapBOToEntity(apartment models.ApartmentBO) models.ApartmentEntity {
	return models.ApartmentEntity{
		ID:         apartment.ID,
		Number:     apartment.Number,
		Coeficient: apartment.Coeficient,
		Meters:     apartment.Meters,
	}

}
func mapEntityToBO(apartment models.ApartmentEntity) *models.ApartmentBO {
	return &models.ApartmentBO{
		ID:         apartment.ID,
		Number:     apartment.Number,
		Coeficient: apartment.Coeficient,
		Meters:     apartment.Meters,
	}
}

func mapEntitiesToBOs(apartments []models.ApartmentEntity) []models.ApartmentBO {
	result := make([]models.ApartmentBO, len(apartments))
	for i, apartment := range apartments {
		result[i] = *mapEntityToBO(apartment)
	}
	return result
}
