package valueobject

import (
	"github.com/pkg/errors"
)

type EmergencyContact struct {
	name  string
	phone *PhoneNumber
}

func NewEmergencyContact(name string, phone *PhoneNumber) (*EmergencyContact, error) {
	if name == "" {
		return nil, errors.New("emergency contact name is required")
	}

	if phone == nil {
		return nil, errors.New("emergency contact phone number is required")
	}

	return &EmergencyContact{
		name:  name,
		phone: phone,
	}, nil
}

func (ec *EmergencyContact) Name() string {
	return ec.name
}

func (ec *EmergencyContact) Phone() *PhoneNumber {
	return ec.phone
}
