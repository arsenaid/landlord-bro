package valueobject

// Address represents a property address.
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
	Country    string
}

// NewAddress creates a new Address value object.
func NewAddress(street, city, state, postalCode, country string) *Address {
	return &Address{
		Street:     street,
		City:       city,
		State:      state,
		PostalCode: postalCode,
		Country:    country,
	}
}
