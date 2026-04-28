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
}
