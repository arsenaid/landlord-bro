package service

import (
	"context"
	"errors"
	"landlord-bro/domain/aggregate"
	"landlord-bro/domain/repository"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

var (
	ErrLeaseNotFound = errors.New("lease not found")
)

type LeaseService struct {
	propertyRepository repository.PropertyRepository
	leaseRepository    repository.LeaseRepository
}

func NewLeaseService(leaseRepo repository.LeaseRepository, propertyRepo repository.PropertyRepository) *LeaseService {
	return &LeaseService{
		leaseRepository:    leaseRepo,
		propertyRepository: propertyRepo,
	}
}

func (ls *LeaseService) CreateLease(
	ctx context.Context,
	propertyID, tenantID uuid.UUID, startDate, endDate time.Time, paymentPeriod aggregate.PaymentPeriod,
) (*aggregate.Lease, error) {
	property, err := ls.propertyForLeaseExists(ctx, propertyID)
	if err != nil {
		return nil, err
	}

	lease, err := property.ReserveProperty(tenantID, startDate, endDate, money.New(100, "USD"), paymentPeriod, propertyID)
	if err != nil {
		return nil, err
	}

	if err := ls.leaseRepository.Create(ctx, lease); err != nil {
		return nil, err
	}

	if err := ls.propertyRepository.Update(ctx, property); err != nil {
		return nil, err
	}

	return lease, nil
}

func (ls *LeaseService) leaseExists(ctx context.Context, leaseID uuid.UUID) (*aggregate.Lease, error) {
	lease, err := ls.leaseRepository.GetByID(ctx, leaseID)
	if err == repository.ErrLeaseNotFound {
		return nil, ErrLeaseNotFound
	}
	return lease, nil
}

func (ls *LeaseService) propertyForLeaseExists(ctx context.Context, propertyID uuid.UUID) (*aggregate.Property, error) {
	property, err := ls.propertyRepository.GetByID(ctx, propertyID)
	if err == repository.ErrPropertyNotFound {
		return nil, ErrPropertyNotFound
	}
	return property, nil
}
