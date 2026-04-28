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
	if err := db.AutoMigrate(&Unit{}, &Owner{}); err != nil {
		return err
	}

	indexes := []string{
		`CREATE UNIQUE INDEX IF NOT EXISTS idx_owner_unique 
         ON owners(identification, unit_id) WHERE deleted_at is NULL`,
		`CREATE UNIQUE INDEX IF NOT EXISTS idx_unit_unique 
         ON units(number, type) WHERE deleted_at is NULL`,
	}

	for _, idx := range indexes {
		if err := db.Exec(idx).Error; err != nil {
			return err
		}
	}
	return nil
}
