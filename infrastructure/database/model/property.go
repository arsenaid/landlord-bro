package model

import (
	"github.com/google/uuid"
)

type Property struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Name         string
	PropertyType PropertyType
	Units        []*PropertyUnit `gorm:"foreignKey:PropertyID"`
}

type PropertyType int

const (
	SingleFamilyHome PropertyType = iota
	MultiUnit
)

type PropertyUnit struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	UnitName   string
	PropertyID uuid.UUID `gorm:"not null"`
	Address    Address   `gorm:"embedded;embeddedPrefix:address_"`
	Rooms      int
	Baths      int
	TotalArea  int
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
	Country    string
}

func (t PropertyType) String() string {
	switch t {
	case SingleFamilyHome:
		return "SingleFamilyHome"
	case MultiUnit:
		return "MultiUnit"
	default:
		return "Unknown"
	}
}
