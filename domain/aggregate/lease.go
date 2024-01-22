package aggregate

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"time"
)

type PaymentPeriod string

const (
	Monthly PaymentPeriod = "monthly"
	Weekly  PaymentPeriod = "weekly"
)

type Lease struct {
	leaseID        uuid.UUID
	propertyUnitID uuid.UUID
	tenantID       uuid.UUID
	startDate      time.Time
	endDate        time.Time
	rentAmount     *money.Money
	paymentPeriod  PaymentPeriod
}

func NewLease(
	tenantID uuid.UUID, propertyID uuid.UUID, startDate, endDate time.Time, rentAmount *money.Money,
	paymentPeriod PaymentPeriod,
) *Lease {
	return &Lease{
		leaseID:        uuid.New(),
		propertyUnitID: propertyID,
		tenantID:       tenantID,
		startDate:      startDate,
		endDate:        endDate,
		rentAmount:     rentAmount,
		paymentPeriod:  paymentPeriod,
	}
}

func (l *Lease) LeaseID() uuid.UUID {
	return l.leaseID
}

func (l *Lease) PropertyUnitID() uuid.UUID {
	return l.propertyUnitID
}

func (l *Lease) TenantID() uuid.UUID {
	return l.tenantID
}

func (l *Lease) StartDate() time.Time {
	return l.startDate
}

func (l *Lease) EndDate() time.Time {
	return l.endDate
}

func (l *Lease) RentAmount() *money.Money {
	return l.rentAmount
}

func (l *Lease) PaymentPeriod() PaymentPeriod {
	return l.paymentPeriod
}
