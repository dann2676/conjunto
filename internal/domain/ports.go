package domain

import (
	"asamblea/internal/models"
	"context"
)

type Gettable[T any] interface {
	GetAll(ctx context.Context) ([]T, error)
}

// Services
type ApartmentService interface {
	Gettable[models.ApartmentBO]
	Get(ctx context.Context, id int) (models.ApartmentBO, error)
	Create(ctx context.Context, apartment models.ApartmentBO) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, apartment models.ApartmentBO) error
}

type OwnerService interface {
	GetAll(ctx context.Context, includeInactive bool) ([]models.OwnerBO, error)
	Get(ctx context.Context, id int) (models.OwnerBO, error)
	Create(ctx context.Context, Owner models.OwnerBO) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, Owner models.OwnerBO) error
}

// Repositories
type ApartmentRepository interface {
	Get(ctx context.Context, id int) (*models.ApartmentBO, error)
	GetAll(ctx context.Context) ([]models.ApartmentBO, error)
	Save(ctx context.Context, apartment models.ApartmentBO) error
	Delete(ctx context.Context, id int) error
}

type OwnerRepository interface {
	Get(ctx context.Context, id int) (*models.OwnerBO, error)
	GetAll(ctx context.Context, includeInactive bool) ([]models.OwnerBO, error)
	Save(ctx context.Context, Owner models.OwnerBO) error
	Delete(ctx context.Context, id int) error
}
