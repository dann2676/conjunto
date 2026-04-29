package models

import (
	"time"

	"gorm.io/gorm"
)

type OwnerEntity struct {
	ID             int        `gorm:"id"`
	Identification string     `gorm:"column:identification"`
	Name           string     `gorm:"column:name"`
	Email          string     `gorm:"column:email"`
	Phone          string     `gorm:"column:phone"`
	StartDate      time.Time  `gorm:"column:start_date"`
	EndDate        *time.Time `gorm:"column:end_date"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	UnitID         int            `gorm:"column:unit_id"`
	Unit           UnitEntity     `gorm:"foreignKey:UnitID"`
}

func (OwnerEntity) TableName() string {
	return "owners"
}

type OwnerBO struct {
	ID             int
	Name           string
	Identification string
	Email          string
	Phone          string
	Unit           int
	UnitID         int
	Active         bool
	StartDate      time.Time
}

type OwnerDTO struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Identification string `json:"identification"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Unit           int    `json:"unit"`
	UnitID         int    `json:"unit_id"`
	Active         bool
}

type OwnerRequest struct {
	Identification string `form:"identification" binding:"required"`
	Name           string `form:"name" binding:"required"`
	Email          string `form:"email" binding:"required,email"`
	Phone          string `form:"phone" binding:"required"`
	UnitID         int    `form:"unit_id" binding:"required"`
}
