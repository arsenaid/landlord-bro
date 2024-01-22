package aggregate

import (
	"fmt"
	"landlord-bro/domain/valueobject"

	"github.com/google/uuid"
)

type TenantBuilder struct {
	tenant *Tenant
}

func NewTenantBuilder() *TenantBuilder {
	return &TenantBuilder{&Tenant{}}
}

func (b *TenantBuilder) WithNewTenantID() *TenantBuilder {
	b.tenant.tenantID = uuid.New()
	return b
}

func (b *TenantBuilder) WithTenantID(tenantID uuid.UUID) *TenantBuilder {
	b.tenant.tenantID = tenantID
	return b
}

func (b *TenantBuilder) WithFirstName(name string) *TenantBuilder {
	b.tenant.firstName = name
	return b
}

func (b *TenantBuilder) WithLastName(name string) *TenantBuilder {
	b.tenant.lastName = name
	return b
}

func (b *TenantBuilder) WithEmail(email string) *TenantBuilder {
	b.tenant.email = email
	return b
}

func (b *TenantBuilder) WithPhoneNumber(phone *valueobject.PhoneNumber) *TenantBuilder {
	b.tenant.phoneNumber = phone
	return b
}

func (b *TenantBuilder) WithEmergencyContact(emergencyContact *valueobject.EmergencyContact) *TenantBuilder {
	b.tenant.emergencyContact = emergencyContact
	return b
}

func (b *TenantBuilder) Build() (*Tenant, error) {
	err := b.tenant.Validate()
	if err != nil {
		return nil, fmt.Errorf("tenant validation failed: %s", err)
	}
	return b.tenant, nil
}
