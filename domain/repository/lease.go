package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"landlord-bro/domain/aggregate"
)

var (
	ErrLeaseNotFound = errors.New("lease not found")
)

type LeaseRepository interface {
	Create(ctx context.Context, lease *aggregate.Lease) error
	Update(ctx context.Context, lease *aggregate.Lease) error
	Delete(ctx context.Context, leaseID uuid.UUID) error
	GetByID(ctx context.Context, leaseID uuid.UUID) (*aggregate.Lease, error)
	GetAll(ctx context.Context, pageToken string, pageSize int32) ([]*aggregate.Lease, string, error)
}
