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
	Slug           string         `gorm:"column:slug"`
}

func (AssemblyEntity) TableName() string { return "assemblies" }

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
	Slug           string
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
	Slug           string
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
	Slug           string
}

type AgendaItemRequest struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description"`
	Order       int    `form:"order" binding:"required,min=1"`
}

type VoteRequest struct {
	Value string `form:"value" binding:"required,oneof=yes no abstain"`
}

type AssemblyUnitEntity struct {
	ID           int          `gorm:"id"`
	AssemblyID   int          `gorm:"column:assembly_id"`
	UnitID       int          `gorm:"column:unit_id"`
	Unit         UnitEntity   `gorm:"foreignKey:UnitID"`
	OwnerID      *int         `gorm:"column:owner_id"`
	Owner        *OwnerEntity `gorm:"foreignKey:OwnerID"`
	AttendedBy   string       `gorm:"column:attended_by"`
	AttendedByID string       `gorm:"column:attended_by_id"`
	IsProxy      bool         `gorm:"column:is_proxy"`
	ProxyFor     string       `gorm:"column:proxy_for"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type AssemblyUnitBO struct {
	ID           int
	AssemblyID   int
	UnitID       int
	UnitNumber   int
	Coeficient   float32
	OwnerID      *int
	AttendedBy   string
	AttendedByID string
	IsProxy      bool
	ProxyFor     string
}

type AttendanceRequest struct {
	Units []AttendanceUnitRequest `form:"units"`
}

type AttendanceUnitRequest struct {
	UnitID   int    `form:"unit_id"`
	IsProxy  bool   `form:"is_proxy"`
	ProxyFor string `form:"proxy_for"`
}

type AttendanceLookupRequest struct {
	Identification string `form:"identification" binding:"required"`
}
