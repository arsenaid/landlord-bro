package translator

import (
	"landlord-bro/domain/aggregate"
	"landlord-bro/domain/valueobject"
	"landlord-bro/infrastructure/database/model"

	"github.com/google/uuid"
)

// ToModel converts a Property domain entity to a Property GORM model.
func ToModel(propertyEntity *aggregate.Property) *model.Property {
	// Convert units if needed
	var units []*model.PropertyUnit
	for _, unitEntity := range propertyEntity.Units() {
		unitModel := toPropertyUnitModel(unitEntity, propertyEntity.PropertyID())
		units = append(units, unitModel)
	}

	return &model.Property{
		ID:           propertyEntity.PropertyID(),
		Name:         propertyEntity.Name(),
		PropertyType: model.PropertyType(propertyEntity.PropertyType()),
		Units:        units,
	}
}

// FromModel converts a Property GORM model to a Property domain entity.
func FromModel(propertyModel *model.Property) *aggregate.Property {
	propertyType := aggregate.PropertyType(propertyModel.PropertyType)

	// Convert units if needed
	var units []*aggregate.PropertyUnit
	for _, unitModel := range propertyModel.Units {
		unitEntity := fromPropertyUnitModel(unitModel)
		units = append(units, unitEntity)
	}

	return aggregate.NewExistingProperty(
		propertyModel.ID,
		propertyModel.Name,
		propertyType,
		units,
	)
}

// ToModel converts a PropertyUnit domain entity to a PropertyUnit GORM model.
func toPropertyUnitModel(propertyUnitEntity *aggregate.PropertyUnit, propertyID uuid.UUID) *model.PropertyUnit {
	return &model.PropertyUnit{
		ID:         propertyUnitEntity.UnitID(),
		UnitName:   propertyUnitEntity.Name(),
		PropertyID: propertyID,
		Address: model.Address{
			Street:     propertyUnitEntity.Address().Street,
			City:       propertyUnitEntity.Address().City,
			State:      propertyUnitEntity.Address().State,
			PostalCode: propertyUnitEntity.Address().PostalCode,
			Country:    propertyUnitEntity.Address().Country,
		},
		Rooms:     propertyUnitEntity.Rooms(),
		Baths:     propertyUnitEntity.Bathrooms(),
		TotalArea: propertyUnitEntity.TotalArea(),
	}
}

// FromModel converts a PropertyUnit GORM model to a PropertyUnit domain entity.
func fromPropertyUnitModel(propertyUnitModel *model.PropertyUnit) *aggregate.PropertyUnit {
	address := valueobject.NewAddress(
		propertyUnitModel.Address.Street,
		propertyUnitModel.Address.City,
		propertyUnitModel.Address.State,
		propertyUnitModel.Address.PostalCode,
		propertyUnitModel.Address.Country,
	)

	return aggregate.NewExistingPropertyUnit(
		propertyUnitModel.ID,
		propertyUnitModel.UnitName,
		address,
		propertyUnitModel.Rooms,
		propertyUnitModel.Baths,
		propertyUnitModel.TotalArea,
	)
}
