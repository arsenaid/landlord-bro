package aggregate

import (
	"fmt"
	"landlord-bro/domain/valueobject"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Tenant represents a tenant in the system.
type Tenant struct {
	tenantID         uuid.UUID
	firstName        string
	lastName         string
	email            string
	phoneNumber      *valueobject.PhoneNumber
	emergencyContact *valueobject.EmergencyContact
	validate         *validator.Validate
}

func (t *Tenant) Validate() error {
	t.validate = validator.New()

	err := t.validate.Var(t.tenantID, "required")
	if err != nil {
		return fmt.Errorf("tenant LeaseID is required")
	}

	err = t.validate.Var(t.firstName, "required")
	if err != nil {
		return fmt.Errorf("first name is required")
	}

	err = t.validate.Var(t.lastName, "required")
	if err != nil {
		return fmt.Errorf("last name is required")
	}

	err = t.validate.Var(t.email, "required,email")
	if err != nil {
		return fmt.Errorf("email is required and should be valid")
	}

	err = t.validate.Struct(t.phoneNumber)
	if err != nil {
		return fmt.Errorf("phone number is required and should be valid")
	}

	return nil
}

func (t *Tenant) TenantID() uuid.UUID {
	return t.tenantID
}

func (t *Tenant) FirstName() string {
	return t.firstName
}

func (t *Tenant) LastName() string {
	return t.lastName
}

func (t *Tenant) Email() string {
	return t.email
}

func (t *Tenant) PhoneNumber() *valueobject.PhoneNumber {
	return t.phoneNumber
}

func (t *Tenant) EmergencyContact() *valueobject.EmergencyContact {
	return t.emergencyContact
}
