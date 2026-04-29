package models

import (
	"time"

	"gorm.io/gorm"
)

// Entities
type AssemblyEntity struct {
	ID             int       `gorm:"id"`
	Title          string    `gorm:"column:title"`
	Date           time.Time `gorm:"column:date"`
	Type           string    `gorm:"column:type"`
	Status         string    `gorm:"column:status"`
	QuorumRequired float32   `gorm:"column:quorum_required"`
	MeetingURL     string    `gorm:"column:meeting_url"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (AssemblyEntity) TableName() string { return "assemblies" }

type AssemblyUnitEntity struct {
	ID            int        `gorm:"id"`
	AssemblyID    int        `gorm:"column:assembly_id"`
	UnitID        int        `gorm:"column:unit_id"`
	Unit          UnitEntity `gorm:"foreignKey:UnitID"`
	AttendedBy    string     `gorm:"column:attended_by"`
	RepresentedBy string     `gorm:"column:represented_by"`
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (AssemblyUnitEntity) TableName() string { return "assembly_units" }

type AgendaItemEntity struct {
	ID          int    `gorm:"id"`
	AssemblyID  int    `gorm:"column:assembly_id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Order       int    `gorm:"column:order"`
	Status      string `gorm:"column:status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (AgendaItemEntity) TableName() string { return "agenda_items" }

type VoteEntity struct {
	ID           int        `gorm:"id"`
	AgendaItemID int        `gorm:"column:agenda_item_id"`
	UnitID       int        `gorm:"column:unit_id"`
	Unit         UnitEntity `gorm:"foreignKey:UnitID"`
	Value        string     `gorm:"column:value"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (VoteEntity) TableName() string { return "votes" }

// BOs
type AssemblyBO struct {
	ID             int
	Title          string
	Date           time.Time
	Type           string
	Status         string
	QuorumRequired float32
	MeetingURL     string
}

type AssemblyUnitBO struct {
	ID            int
	AssemblyID    int
	UnitID        int
	UnitNumber    int
	Coeficient    float32
	AttendedBy    string
	RepresentedBy string
	IsProxy       bool // true si RepresentedBy no está vacío
}

type AgendaItemBO struct {
	ID          int
	AssemblyID  int
	Title       string
	Description string
	Order       int
	Status      string
}

type VoteBO struct {
	ID           int
	AgendaItemID int
	UnitID       int
	UnitNumber   int
	Coeficient   float32
	Value        string
}

// DTOs
type AssemblyDTO struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Date           time.Time `json:"date"`
	Type           string    `json:"type"`
	Status         string    `json:"status"`
	QuorumRequired float32   `json:"quorum_required"`
	MeetingURL     string    `json:"meeting_url"`
}

type AssemblyUnitDTO struct {
	ID            int     `json:"id"`
	UnitNumber    int     `json:"unit_number"`
	Coeficient    float32 `json:"coeficient"`
	AttendedBy    string  `json:"attended_by"`
	RepresentedBy string  `json:"represented_by"`
	IsProxy       bool    `json:"is_proxy"`
}

type AgendaItemDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	Status      string `json:"status"`
}

type VoteDTO struct {
	ID         int     `json:"id"`
	UnitNumber int     `json:"unit_number"`
	Coeficient float32 `json:"coeficient"`
	Value      string  `json:"value"`
}

// Requests
type AssemblyRequest struct {
	Title          string  `form:"title" binding:"required"`
	Date           string  `form:"date" binding:"required"`
	Type           string  `form:"type" binding:"required,oneof=ordinaria extraordinaria"`
	QuorumRequired float32 `form:"quorum_required" binding:"required,gt=0,lte=1"`
	MeetingURL     string  `form:"meeting_url"`
}

type AgendaItemRequest struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description"`
	Order       int    `form:"order" binding:"required,min=1"`
}

type AttendanceRequest struct {
	UnitID        int    `form:"unit_id" binding:"required"`
	AttendedBy    string `form:"attended_by" binding:"required"`
	RepresentedBy string `form:"represented_by"`
}

type VoteRequest struct {
	Value string `form:"value" binding:"required,oneof=yes no abstain"`
}
