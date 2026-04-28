package unit

import "asamblea/internal/models"

func mapBOToEntity(unit models.UnitBO) models.UnitEntity {
	return models.UnitEntity{
		ID:         unit.ID,
		Number:     unit.Number,
		Coeficient: unit.Coeficient,
		Meters:     unit.Meters,
		Type:       unit.Type,
	}

}
func mapEntityToBO(unit models.UnitEntity) models.UnitBO {
	return models.UnitBO{
		ID:         unit.ID,
		Number:     unit.Number,
		Coeficient: unit.Coeficient,
		Meters:     unit.Meters,
		Type:       unit.Type,
	}
}

func mapEntitiesToBOs(units []models.UnitEntity) []models.UnitBO {
	result := make([]models.UnitBO, len(units))
	for i, unit := range units {
		result[i] = mapEntityToBO(unit)
	}
	return result
}
