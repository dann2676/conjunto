package assembly

import (
	"asamblea/internal/models"
	"time"
)

func mapRequestToBO(r models.AssemblyRequest) (models.AssemblyBO, error) {
	date, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		return models.AssemblyBO{}, err
	}
	return models.AssemblyBO{
		Title:          r.Title,
		Date:           date,
		Type:           r.Type,
		QuorumRequired: r.QuorumRequired,
		MeetingURL:     r.MeetingURL,
	}, nil
}

func mapBOToDTO(a models.AssemblyBO) models.AssemblyDTO {
	return models.AssemblyDTO{
		ID:             a.ID,
		Title:          a.Title,
		Date:           a.Date,
		Type:           a.Type,
		Status:         a.Status,
		QuorumRequired: a.QuorumRequired,
		MeetingURL:     a.MeetingURL,
		Slug:           a.Slug,
	}
}

func mapBOsToDTOs(as []models.AssemblyBO) []models.AssemblyDTO {
	result := make([]models.AssemblyDTO, len(as))
	for i, a := range as {
		result[i] = mapBOToDTO(a)
	}
	return result
}

func mapAgendaRequestToBO(r models.AgendaItemRequest, assemblyID int) models.AgendaItemBO {
	return models.AgendaItemBO{
		AssemblyID:  assemblyID,
		Title:       r.Title,
		Description: r.Description,
		Order:       r.Order,
	}
}

func mapAgendaBOToDTO(item models.AgendaItemBO) models.AgendaItemDTO {
	return models.AgendaItemDTO{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Order:       item.Order,
		Status:      item.Status,
	}
}

func mapAgendaBOsToDTOs(items []models.AgendaItemBO) []models.AgendaItemDTO {
	result := make([]models.AgendaItemDTO, len(items))
	for i, item := range items {
		result[i] = mapAgendaBOToDTO(item)
	}
	return result
}
