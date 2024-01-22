package repository

import (
	"context"
	"landlord-bro/domain/aggregate"
	"landlord-bro/domain/repository"

	"landlord-bro/infrastructure/database"
	"landlord-bro/infrastructure/database/model"
	"landlord-bro/infrastructure/database/translator"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PropertyRepository struct {
	db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) *PropertyRepository {
	return &PropertyRepository{db}
}

func (r *PropertyRepository) Create(ctx context.Context, property *aggregate.Property) error {
	model := translator.ToModel(property)
	return r.db.Create(model).Error
}

func (r *PropertyRepository) Update(ctx context.Context, property *aggregate.Property) error {
	model := translator.ToModel(property)
	return r.db.Save(model).Error
}

func (r *PropertyRepository) Delete(ctx context.Context, propertyID uuid.UUID) error {
	return r.db.Where("id = ?", propertyID).Delete(&model.Property{}).Error
}

func (r *PropertyRepository) GetByID(ctx context.Context, propertyID uuid.UUID) (*aggregate.Property, error) {
	var model model.Property

	err := r.db.Preload("Units").Where("id = ?", propertyID).First(&model).Error
	if err == gorm.ErrRecordNotFound {
		return nil, repository.ErrPropertyNotFound
	}

	if err != nil {
		return nil, err
	}
	return translator.FromModel(&model), nil
}

func (r *PropertyRepository) GetAll(ctx context.Context, pageToken string, pageSize int32) ([]*aggregate.Property, string, error) {
	var properties []*aggregate.Property
	var models []*model.Property

	query := r.db

	// Apply eager loading for the Units association
	query = query.Preload("Units")

	// Define the base query to select properties
	query = query.Model(&model.Property{})

	if pageToken != "" {
		pageUUID, err := database.DecodeCursor(pageToken)
		if err != nil {
			return nil, "", err
		}
		var lastProperty model.Property
		if err := query.Where("id = ?", pageUUID).First(&lastProperty).Error; err != nil {
			return nil, "", err
		}
		query = query.Where("id > ?", pageToken)
	}

	// Limit the result set by the specified page size
	if pageSize > 0 {
		query = query.Limit(int(pageSize))
	}

	// Execute the query to retrieve properties
	if err := query.Find(&models).Error; err != nil {
		return nil, "", err
	}

	// Determine the next page token based on the last property's ID
	nextPageToken := ""
	if len(properties) > 0 {
		lastProperty := models[len(properties)-1]
		nextPageToken = database.EncodeCursor(lastProperty.ID)
	}

	for _, model := range models {
		property := translator.FromModel(model)
		properties = append(properties, property)
	}

	return properties, nextPageToken, nil
}
