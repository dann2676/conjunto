package models

import (
	"time"

	"gorm.io/gorm"
)

type UnitEntity struct {
	ID         int     `gorm:"id"`
	Number     int     `gorm:"column:number"`
	Coeficient float32 `gorm:"column:coeficient"`
	Meters     float32 `gorm:"column:meters"`
	Type       string  `gorm:"column:type"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (UnitEntity) TableName() string {
	return "units"
}

type UnitBO struct {
	ID         int
	Number     int
	Coeficient float32
	Meters     float32
	Type       string
}

type UnitDTO struct {
	ID         int     `json:"id"`
	Number     int     `json:"number"`
	Coeficient float32 `json:"coeficient"`
	Meters     float32 `json:"meters"`
	Type       string  `json:"type"`
}

type UnitRequest struct {
	Number     int     `form:"number" binding:"required,min=10,max=999"`
	Coeficient float32 `form:"coeficient" binding:"required,gt=0"`
	Meters     float32 `form:"meters" binding:"required,gt=1"`
	Type       string  `form:"type" binding:"required,oneof=apartment parking"`
}
