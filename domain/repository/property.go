package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"landlord-bro/domain/aggregate"
)

var (
	ErrPropertyNotFound = errors.New("property not found")
)

type PropertyRepository interface {
	Create(ctx context.Context, property *aggregate.Property) error
	Update(ctx context.Context, property *aggregate.Property) error
	Delete(ctx context.Context, propertyID uuid.UUID) error
	GetByID(ctx context.Context, propertyID uuid.UUID) (*aggregate.Property, error)
	GetAll(ctx context.Context, pageToken string, pageSize int32) ([]*aggregate.Property, string, error)
}
