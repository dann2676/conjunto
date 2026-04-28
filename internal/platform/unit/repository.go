package unit

import (
	"asamblea/internal/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UnitRepository {
	return &repository{db: db}
}
