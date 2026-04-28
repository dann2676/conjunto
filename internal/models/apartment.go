package models

import "time"

type ApartmentEntity struct {
	ID         int     `gorm:"id"`
	Number     int     `gorm:"column:number"`
	Coeficient float32 `gorm:"column:coeficient"`
	Meters     float32 `gorm:"column:meters"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func (ApartmentEntity) TableName() string {
	return "apartments"
}

type ApartmentBO struct {
	ID         int
	Number     int
	Coeficient float32
	Meters     float32
}

type ApartmentDTO struct {
	ID         int     `json:"id"`
	Number     int     `json:"number"`
	Coeficient float32 `json:"coeficient"`
	Meters     float32 `json:"meters"`
}

type ApartmentRequest struct {
	Number     int     `form:"number" binding:"required,min=100,max=999"`
	Coeficient float32 `form:"coeficient" binding:"required,gt=0"`
	Meters     float32 `form:"meters" binding:"required,gt=1"`
}
