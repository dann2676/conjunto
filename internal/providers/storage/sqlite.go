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

type Apartment struct {
	gorm.Model
	Number     uint    `gorm:"uniqueIndex;not null"`
	Meters     float32 `gorm:"type:real"`
	Coeficient float32 `gorm:"type:real"`
}

type Owner struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Email       string
	Phone       string
	ApartmentID uint
	Active      bool
	Apartment   Apartment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	StartDate   time.Time  `gorm:"type:date"`
	EndDate     *time.Time `gorm:"type:date"`
}

func Init() (*gorm.DB, error) {
	dns := dbName + "?_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_busy_timeout=5000"
	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{
		TranslateError: true,
	})
	db.AutoMigrate(&Apartment{}, &Owner{})

	db.Exec("PRAGMA foreign_keys = ON;")
	return db, err
}
