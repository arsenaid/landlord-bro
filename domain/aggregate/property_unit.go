package aggregate

import (
	"landlord-bro/domain/valueobject"

	"github.com/google/uuid"
)

type PropertyUnit struct {
	unitID    uuid.UUID
	name      string
	address   *valueobject.Address
	rooms     int
	bathrooms int
	totalArea int
}

func NewPropertyUnit(name string, address *valueobject.Address, rooms int, bathrooms int, totalArea int) *PropertyUnit {
	return &PropertyUnit{
		unitID:    uuid.New(),
		name:      name,
		address:   address,
		rooms:     rooms,
		bathrooms: bathrooms,
		totalArea: totalArea,
	}
}

func NewExistingPropertyUnit(
	unitID uuid.UUID, name string, address *valueobject.Address, rooms int, bathrooms int,
	totalArea int,
) *PropertyUnit {
	return &PropertyUnit{
		unitID:    unitID,
		name:      name,
		address:   address,
		rooms:     rooms,
		bathrooms: bathrooms,
		totalArea: totalArea,
	}
}

func (pu *PropertyUnit) UnitID() uuid.UUID {
	return pu.unitID
}

func (pu *PropertyUnit) Name() string {
	return pu.name
}

func (pu *PropertyUnit) Address() *valueobject.Address {
	return pu.address
}

func (pu *PropertyUnit) Rooms() int {
	return pu.rooms
}

func (pu *PropertyUnit) Bathrooms() int {
	return pu.bathrooms
}

func (pu *PropertyUnit) TotalArea() int {
	return pu.totalArea
}
