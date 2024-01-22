package valueobject

import (
	"errors"
	"testing"
)

func TestPhoneNumberEquals(t *testing.T) {
	phoneNumber, err := NewPhoneNumber("1", "555", "1234567")
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	otherPhoneNumber, err := NewPhoneNumber("1", "555", "1234567")
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if !phoneNumber.Equals(otherPhoneNumber) {
		t.Error("Expected phone numbers to be equal")
	}
}

func TestPhoneNumberFormat(t *testing.T) {
	phoneNumber, err := NewPhoneNumber("1", "555", "1234567")
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	expected := "+1 (555) 123-4567"
	actual := phoneNumber.Format()
	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}

func TestPhoneNumberValidate(t *testing.T) {
	_, actual := NewPhoneNumber("1", "555", "")

	expected := errors.New("phone number validation failed: [Number is required]")
	if actual.Error() != expected.Error() {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
