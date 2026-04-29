package assembly

import (
	"asamblea/internal/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.AssemblyRepository {
	return &repository{db: db}
}
