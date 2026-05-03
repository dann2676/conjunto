package domain

import (
	"asamblea/internal/models"
	"context"
)

type Gettable[T any] interface {
	GetAll(ctx context.Context, includeInactive bool) ([]T, error)
	Get(ctx context.Context, id int) (T, error)
}

type Deletable interface {
	Delete(ctx context.Context, id int) error
	Purge(ctx context.Context, id int) error
}

// Services
type UnitService interface {
	Gettable[models.UnitBO]
	Deletable
	Create(ctx context.Context, unit models.UnitBO) error
	Update(ctx context.Context, unit models.UnitBO) error
}

type OwnerService interface {
	Deletable
	Gettable[models.OwnerBO]
	Create(ctx context.Context, Owner models.OwnerBO) error
	Update(ctx context.Context, Owner models.OwnerBO) error
	GetActiveByUnit(ctx context.Context, unitID int) (models.OwnerBO, error)
	GetByIdentification(ctx context.Context, identification string) (models.OwnerBO, error)
	GetUnitsByOwner(ctx context.Context, ownerID int) ([]models.UnitBO, error)
}

// Repositories
type UnitRepository interface {
	Deletable
	Gettable[models.UnitBO]
	Save(ctx context.Context, unit models.UnitBO) error
}

type OwnerRepository interface {
	Deletable
	Gettable[models.OwnerBO]
	Save(ctx context.Context, Owner models.OwnerBO) error
	GetActiveByUnit(ctx context.Context, unitID int) (models.OwnerBO, error)
	GetByIdentification(ctx context.Context, identification string) (models.OwnerBO, error)
	GetUnitsByOwner(ctx context.Context, ownerID int) ([]models.UnitBO, error)
}

type AssemblyService interface {
	Gettable[models.AssemblyBO]
	Create(ctx context.Context, assembly models.AssemblyBO) error
	Update(ctx context.Context, assembly models.AssemblyBO) error
	Delete(ctx context.Context, id int) error
	UpdateStatus(ctx context.Context, id int, status string) error
	GetBySlug(ctx context.Context, slug string) (models.AssemblyBO, error)

	// Agenda
	GetAgendaItems(ctx context.Context, assemblyID int) ([]models.AgendaItemBO, error)
	CreateAgendaItem(ctx context.Context, item models.AgendaItemBO) error
	UpdateAgendaItemStatus(ctx context.Context, itemID int, status string) error
	DeleteAgendaItem(ctx context.Context, itemID int) error

	// Asistencia y quórum
	RegisterAttendance(ctx context.Context, attendance models.AssemblyUnitBO) error
	GetAttendance(ctx context.Context, assemblyID int) ([]models.AssemblyUnitBO, error)
	GetQuorum(ctx context.Context, assemblyID int) (float32, error)

	// Votaciones
	RegisterVote(ctx context.Context, vote models.VoteBO) error
	GetVotes(ctx context.Context, agendaItemID int) ([]models.VoteBO, error)
}

type AssemblyRepository interface {
	Gettable[models.AssemblyBO]
	Save(ctx context.Context, assembly models.AssemblyBO) error
	Delete(ctx context.Context, id int) error
	GetAgendaItem(ctx context.Context, itemID int) (models.AgendaItemBO, error)
	GetAgendaItems(ctx context.Context, assemblyID int) ([]models.AgendaItemBO, error)
	SaveAgendaItem(ctx context.Context, item models.AgendaItemBO) error
	DeleteAgendaItem(ctx context.Context, itemID int) error
	GetBySlug(ctx context.Context, slug string) (models.AssemblyBO, error)

	RegisterAttendance(ctx context.Context, attendance models.AssemblyUnitBO) error
	GetAttendance(ctx context.Context, assemblyID int) ([]models.AssemblyUnitBO, error)

	RegisterVote(ctx context.Context, vote models.VoteBO) error
	GetVotes(ctx context.Context, agendaItemID int) ([]models.VoteBO, error)
}
