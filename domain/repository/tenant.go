package repository

import (
	"context"
	"landlord-bro/domain/aggregate"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrTenantNotFound = errors.New("tenant not found")
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *aggregate.Tenant) error
	Update(ctx context.Context, tenant *aggregate.Tenant) error
	Delete(ctx context.Context, tenantID uuid.UUID) error
	GetByID(ctx context.Context, tenantID uuid.UUID) (*aggregate.Tenant, error)
	GetAll(ctx context.Context, pageToken string, pageSize int32) ([]*aggregate.Tenant, string, error)
}
