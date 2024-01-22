package service

import (
	"context"
	"errors"
	"landlord-bro/domain/aggregate"
	"landlord-bro/domain/repository"

	"github.com/google/uuid"
)

var (
	ErrTenantNotFound = errors.New("tenant not found")
)

type TenantService struct {
	tenantRepository repository.TenantRepository
}

func NewTenantService(tenantRepo repository.TenantRepository) *TenantService {
	return &TenantService{
		tenantRepository: tenantRepo,
	}
}

func (t *TenantService) CreateTenant(ctx context.Context, tenant *aggregate.Tenant) error {
	// Validate property data
	//if err := property.Validate(); err != nil {
	//	return err
	//}

	return t.tenantRepository.Create(ctx, tenant)
}

func (t *TenantService) UpdateTenant(ctx context.Context, tenant *aggregate.Tenant) error {
	// Validate property data
	//if err := tenant.Validate(); err != nil {
	//	return err
	//}

	if _, err := t.checkTenantNotFound(ctx, tenant.TenantID()); err != nil {
		return err
	}
	return t.tenantRepository.Update(ctx, tenant)
}

func (t *TenantService) DeleteTenant(ctx context.Context, propertyID uuid.UUID) error {
	if _, err := t.checkTenantNotFound(ctx, propertyID); err != nil {
		return err
	}
	return t.tenantRepository.Delete(ctx, propertyID)
}

func (t *TenantService) GetTenantByID(ctx context.Context, tenantID uuid.UUID) (*aggregate.Tenant, error) {
	return t.checkTenantNotFound(ctx, tenantID)
}

func (t *TenantService) ListTenants(ctx context.Context, pageToken string, pageSize int32) ([]*aggregate.Tenant, string, error) {
	return t.tenantRepository.GetAll(ctx, pageToken, pageSize)
}

func (t *TenantService) checkTenantNotFound(ctx context.Context, tenantID uuid.UUID) (*aggregate.Tenant, error) {
	property, err := t.tenantRepository.GetByID(ctx, tenantID)
	if err == repository.ErrTenantNotFound {
		return nil, ErrTenantNotFound
	}
	return property, nil
}
