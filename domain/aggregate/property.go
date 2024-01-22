package aggregate

import (
	"errors"
	"landlord-bro/domain/valueobject"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type PropertyType int

const (
	SingleFamilyHome PropertyType = iota
	MultiUnit
)

type Property struct {
	propertyID   uuid.UUID
	name         string
	propertyType PropertyType
	units        []*PropertyUnit
	leases       []*Lease
}

// NewProperty creates a new Property aggregate.
func NewProperty(
	name string, propertyType PropertyType, units []*PropertyUnit,
) (*Property, error) {
	if name == "" {
		return nil, errors.New("property name is required")
	}
	if len(units) == 0 {
		return nil, errors.New("at least one rental unit is required")
	}

	return &Property{
		propertyID:   uuid.New(),
		name:         name,
		propertyType: propertyType,
		units:        units,
	}, nil
}

// NewExistingProperty creates an existing Property aggregate.
func NewExistingProperty(
	propertyID uuid.UUID, name string, propertyType PropertyType, units []*PropertyUnit,
) *Property {
	property := &Property{
		propertyID:   propertyID,
		name:         name,
		propertyType: propertyType,
		units:        units,
	}

	return property
}

// AddUnit adds a new unit to the RentalProperty.
func (p *Property) AddUnit(unitName string, address *valueobject.Address, rooms, baths, totalArea int) {
	unit := NewPropertyUnit(unitName, address, rooms, baths, totalArea)
	p.units = append(p.units, unit)
}

// ReserveProperty reserves a property unit with a lease.
func (p *Property) ReserveProperty(
	tenantID uuid.UUID, startDate, endDate time.Time, rentAmount *money.Money, paymentPeriod PaymentPeriod, unitID uuid.UUID,
) (*Lease, error) {
	if p.isReserved(startDate, endDate) {
		return nil, errors.New("property unit is already reserved for the given dates")
	}

	lease := NewLease(tenantID, unitID, startDate, endDate, rentAmount, paymentPeriod)
	p.leases = append(p.leases, lease)

	return lease, nil
}

// isReserved checks if the property unit is already reserved for the given dates.
func (p *Property) isReserved(startDate, endDate time.Time) bool {
	for _, lease := range p.leases {
		if startDate.Before(lease.EndDate()) && endDate.After(lease.StartDate()) {
			return true
		}
	}
	return false
}

func (p *Property) PropertyID() uuid.UUID {
	return p.propertyID
}

func (p *Property) Name() string {
	return p.name
}

func (p *Property) SetName(name string) {
	p.name = name
}

func (p *Property) PropertyType() PropertyType {
	return p.propertyType
}

func (p *Property) Units() []*PropertyUnit {
	return p.units
}
