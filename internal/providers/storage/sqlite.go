package storage

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "conjunto.db"

type storage struct {
	*gorm.DB
}

type Unit struct {
	gorm.Model
	Number     uint
	Meters     float32 `gorm:"type:real"`
	Coeficient float32 `gorm:"type:real"`
	Type       string
}

type Owner struct {
	gorm.Model
	Name           string `gorm:"not null"`
	Identification string `gorm:"not null"`
	Email          string
	Phone          string
	UnitID         uint
	Unit           Unit       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	StartDate      time.Time  `gorm:"type:date"`
	EndDate        *time.Time `gorm:"type:date"`
}

type Assembly struct {
	gorm.Model
	Title          string
	Date           time.Time `gorm:"type:date"`
	Type           string    // ordinaria | extraordinaria
	Status         string    // draft | open | closed
	QuorumRequired float32   `gorm:"default:0.5"`
	MeetingURL     string    // opcional, para asambleas virtuales (3g)
	Slug           string    `gorm:"uniqueIndex;not null"`
}

type AssemblyUnit struct {
	gorm.Model
	AssemblyID   uint
	Assembly     Assembly `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UnitID       uint
	Unit         Unit   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	OwnerID      *uint  // nullable — nil si es apoderado externo
	Owner        *Owner `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	AttendedBy   string // nombre de quien asiste físicamente
	AttendedByID string // cédula de quien asiste
	IsProxy      bool   // true si asiste como apoderado
	ProxyFor     string // nombre del propietario que representa si IsProxy
}

type AgendaItem struct {
	gorm.Model
	AssemblyID  uint
	Assembly    Assembly `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Title       string
	Description string
	Order       int
	Status      string // pending | open | closed
}

type Vote struct {
	gorm.Model
	AgendaItemID uint
	AgendaItem   AgendaItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UnitID       uint
	Unit         Unit   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Value        string // yes | no | abstain
}
type AssemblyCode struct {
	gorm.Model
	AssemblyID uint
	UnitID     uint
	Code       string `gorm:"uniqueIndex"`
	Used       bool
}

func Init() (*gorm.DB, error) {
	dns := dbName + "?_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_busy_timeout=5000"
	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{
		TranslateError: true,
	})

	migrate(db)

	db.Exec("PRAGMA foreign_keys = ON;")
	return db, err
}

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&Unit{}, &Owner{}, &Assembly{}, &AssemblyUnit{}, &AgendaItem{}, &Vote{}); err != nil {
		return err
	}

	indexes := []string{
		`CREATE UNIQUE INDEX IF NOT EXISTS idx_owner_unique 
         ON owners(identification, unit_id) WHERE deleted_at is NULL`,
		`CREATE UNIQUE INDEX IF NOT EXISTS idx_unit_unique 
         ON units(number, type) WHERE deleted_at is NULL`,
		`CREATE UNIQUE INDEX IF NOT EXISTS idx_vote_unique 
 ON votes(agenda_item_id, unit_id) WHERE deleted_at IS NULL`,
		`CREATE UNIQUE INDEX IF NOT EXISTS idx_assembly_unit_unique 
 ON assembly_units(assembly_id, unit_id) WHERE deleted_at IS NULL`,
	}

	for _, idx := range indexes {
		if err := db.Exec(idx).Error; err != nil {
			return err
		}
	}
	return nil
}
