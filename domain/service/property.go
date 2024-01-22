package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"landlord-bro/domain/aggregate"
	"landlord-bro/domain/repository"
)

var (
	ErrPropertyNotFound = errors.New("property not found")
)

type PropertyService struct {
	propertyRepository repository.PropertyRepository
}

func NewPropertyService(propertyRepo repository.PropertyRepository) *PropertyService {
	return &PropertyService{
		propertyRepository: propertyRepo,
	}
}

func (s *PropertyService) CreateProperty(ctx context.Context, property *aggregate.Property) error {
	// Validate property data
	//if err := property.Validate(); err != nil {
	//	return err
	//}

	return s.propertyRepository.Create(ctx, property)
}

func (s *PropertyService) UpdateProperty(ctx context.Context, property *aggregate.Property) error {
	// Validate property data
	//if err := tenant.Validate(); err != nil {
	//	return err
	//}

	if _, err := s.propertyExists(ctx, property.PropertyID()); err != nil {
		return err
	}
	return s.propertyRepository.Update(ctx, property)
}

func (s *PropertyService) DeleteProperty(ctx context.Context, propertyID uuid.UUID) error {
	if _, err := s.propertyExists(ctx, propertyID); err != nil {
		return err
	}
	return s.propertyRepository.Delete(ctx, propertyID)
}

func (s *PropertyService) GetPropertyByID(ctx context.Context, propertyID uuid.UUID) (*aggregate.Property, error) {
	return s.propertyExists(ctx, propertyID)
}

func (s *PropertyService) ListProperties(ctx context.Context, pageToken string, pageSize int32) ([]*aggregate.Property, string, error) {
	return s.propertyRepository.GetAll(ctx, pageToken, pageSize)
}

func (s *PropertyService) propertyExists(ctx context.Context, propertyID uuid.UUID) (*aggregate.Property, error) {
	property, err := s.propertyRepository.GetByID(ctx, propertyID)
	if err == repository.ErrPropertyNotFound {
		return nil, ErrPropertyNotFound
	}
	return property, nil
}
