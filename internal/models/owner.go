package models

import "time"

type OwnerEntity struct {
	ID          int        `gorm:"id"`
	Name        string     `gorm:"column:name"`
	Email       string     `gorm:"column:email"`
	Phone       string     `gorm:"column:phone"`
	StartDate   time.Time  `gorm:"column:start_date"`
	EndDate     *time.Time `gorm:"column:end_date"`
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	ApartmentID int             `gorm:"column:apartment_id"`
	Apartment   ApartmentEntity `gorm:"foreignKey:ApartmentID"`
}

func (OwnerEntity) TableName() string {
	return "owners"
}

type OwnerBO struct {
	ID          int
	Name        string
	Email       string
	Phone       string
	Apartment   int
	ApartmentID int
}

type OwnerDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Apartment   int    `json:"apartment"`
	ApartmentID int    `json:"apartment_id"`
}

type OwnerRequest struct {
	Name        string `form:"name" binding:"required"`
	Email       string `form:"email" binding:"required,email"`
	Phone       string `form:"phone" binding:"required"`
	ApartmentID int    `form:"apartment_id" binding:"required"`
}
