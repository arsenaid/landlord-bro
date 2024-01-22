package valueobject

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type PhoneNumber struct {
	CountryCode string `validate:"required,numeric,len=1"`
	AreaCode    string `validate:"required,numeric,len=3"`
	Number      string `validate:"required,numeric,len=7"`
	validate    *validator.Validate
}

func NewPhoneNumber(countryCode, areaCode, number string) (*PhoneNumber, error) {
	p := &PhoneNumber{
		CountryCode: countryCode,
		AreaCode:    areaCode,
		Number:      number,
	}

	p.validate = validator.New()

	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *PhoneNumber) Validate() error {
	if err := p.validate.Struct(p); err != nil {
		var errorMessages []string

		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}

		return fmt.Errorf("phone number validation failed: %s", errorMessages)
	}

	return nil
}

func (p *PhoneNumber) Format() string {
	return fmt.Sprintf(
		"+%s (%s) %s-%s",
		p.CountryCode,
		p.AreaCode,
		p.Number[:3],
		p.Number[3:],
	)
}

func (p *PhoneNumber) Equals(other *PhoneNumber) bool {
	// compare this phone number to another phone number for equality
	if other == nil {
		return false
	}
	return p.CountryCode == other.CountryCode &&
		p.AreaCode == other.AreaCode &&
		p.Number == other.Number
}
