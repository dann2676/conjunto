package assembly

import (
	"asamblea/internal/models"
)

func mapAssemblyToEntity(a models.AssemblyBO) models.AssemblyEntity {
	return models.AssemblyEntity{
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

func mapEntityToAssembly(a models.AssemblyEntity) models.AssemblyBO {
	return models.AssemblyBO{
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

func mapEntitiesToAssemblies(as []models.AssemblyEntity) []models.AssemblyBO {
	result := make([]models.AssemblyBO, len(as))
	for i, a := range as {
		result[i] = mapEntityToAssembly(a)
	}
	return result
}

func mapAgendaItemToEntity(item models.AgendaItemBO) models.AgendaItemEntity {
	return models.AgendaItemEntity{
		ID:          item.ID,
		AssemblyID:  item.AssemblyID,
		Title:       item.Title,
		Description: item.Description,
		Order:       item.Order,
		Status:      item.Status,
	}
}

func mapEntityToAgendaItem(item models.AgendaItemEntity) models.AgendaItemBO {
	return models.AgendaItemBO{
		ID:          item.ID,
		AssemblyID:  item.AssemblyID,
		Title:       item.Title,
		Description: item.Description,
		Order:       item.Order,
		Status:      item.Status,
	}
}

func mapAttendanceToEntity(a models.AssemblyUnitBO) models.AssemblyUnitEntity {
	return models.AssemblyUnitEntity{
		ID:           a.ID,
		AssemblyID:   a.AssemblyID,
		UnitID:       a.UnitID,
		OwnerID:      a.OwnerID,
		AttendedBy:   a.AttendedBy,
		AttendedByID: a.AttendedByID,
		IsProxy:      a.IsProxy,
		ProxyFor:     a.ProxyFor,
	}
}

func mapEntityToAttendance(a models.AssemblyUnitEntity) models.AssemblyUnitBO {
	return models.AssemblyUnitBO{
		ID:           a.ID,
		AssemblyID:   a.AssemblyID,
		UnitID:       a.UnitID,
		UnitNumber:   a.Unit.Number,
		Coeficient:   a.Unit.Coeficient,
		OwnerID:      a.OwnerID,
		AttendedBy:   a.AttendedBy,
		AttendedByID: a.AttendedByID,
		IsProxy:      a.IsProxy,
		ProxyFor:     a.ProxyFor,
	}
}

func mapVoteToEntity(v models.VoteBO) models.VoteEntity {
	return models.VoteEntity{
		ID:           v.ID,
		AgendaItemID: v.AgendaItemID,
		UnitID:       v.UnitID,
		Value:        v.Value,
	}
}

func mapEntityToVote(v models.VoteEntity) models.VoteBO {
	return models.VoteBO{
		ID:           v.ID,
		AgendaItemID: v.AgendaItemID,
		UnitID:       v.UnitID,
		UnitNumber:   v.Unit.Number,
		Coeficient:   v.Unit.Coeficient,
		Value:        v.Value,
	}
}
