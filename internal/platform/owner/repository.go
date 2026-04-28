package owner

import (
	"asamblea/internal/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.OwnerRepository {
	return &repository{db: db}
}
