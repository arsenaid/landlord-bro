package database

import (
	"github.com/google/uuid"
)

// DecodeCursor Helper function to decode the cursor
func DecodeCursor(cursor string) (*uuid.UUID, error) {
	decodedCursor, err := uuid.Parse(cursor)
	if err != nil {
		return &uuid.UUID{}, err
	}
	return &decodedCursor, nil
}

// EncodeCursor Helper function to encode the cursor
func EncodeCursor(cursor uuid.UUID) string {
	return cursor.String()
}
