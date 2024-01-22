package translator

import (
	"errors"
	"landlord-bro/domain/aggregate"
	"landlord-bro/domain/valueobject"
	proto "landlord-bro/interfaces/api/gen/v1"

	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// FromProtoToEntity converts a RentalProperty proto message to a Property domain entity.
func FromProtoToEntity(protoProperty *proto.RentalProperty) (*aggregate.Property, error) {
	propertyType, err := ProtoPropertyTypeToDomain(protoProperty.GetPropertyType())
	if err != nil {
		return nil, err
	}

	units := make([]*aggregate.PropertyUnit, len(protoProperty.GetUnits()))
	for i, protoUnit := range protoProperty.GetUnits() {
		unit := aggregate.NewPropertyUnit(
			protoUnit.GetUnitName(),
			valueobject.NewAddress(
				protoUnit.GetAddress().GetStreet(),
				protoUnit.GetAddress().GetCity(),
				protoUnit.GetAddress().GetState(),
				protoUnit.GetAddress().GetPostalCode(),
				protoUnit.GetAddress().GetCountry(),
			),
			int(protoUnit.GetBeds()),
			int(protoUnit.GetBaths()),
			int(protoUnit.GetTotalArea()),
		)
		units[i] = unit
	}

	property, err := aggregate.NewProperty(
		protoProperty.GetPropertyName(),
		propertyType,
		units,
	)

	if err != nil {
		return nil, err
	}

	return property, nil
}

// ToProtoFromEntity converts a Property domain entity to a RentalProperty proto message.
func ToProtoFromEntity(entityProperty *aggregate.Property) *proto.RentalProperty {
	protoUnits := make([]*proto.PropertyUnit, len(entityProperty.Units()))
	for i, unit := range entityProperty.Units() {
		protoUnits[i] = &proto.PropertyUnit{
			UnitId:   unit.UnitID().String(),
			UnitName: unit.Name(),
			Address: &proto.Address{
				Street:     unit.Address().Street,
				City:       unit.Address().City,
				State:      unit.Address().State,
				PostalCode: unit.Address().PostalCode,
				Country:    unit.Address().Country,
			},
			Beds:      int32(unit.Rooms()),
			Baths:     int32(unit.Bathrooms()),
			TotalArea: int32(unit.TotalArea()),
		}
	}

	return &proto.RentalProperty{
		PropertyId:   entityProperty.PropertyID().String(),
		PropertyName: entityProperty.Name(),
		PropertyType: DomainPropertyTypeToProto(entityProperty.PropertyType()),
		Units:        protoUnits,
	}
}

// ProtoPropertyTypeToDomain converts a PropertyType enum from proto to the domain entity.
func ProtoPropertyTypeToDomain(protoType proto.PropertyType) (aggregate.PropertyType, error) {
	switch protoType {
	case proto.PropertyType_SINGLE_FAMILY_HOME:
		return aggregate.SingleFamilyHome, nil
	case proto.PropertyType_MULTI_UNIT:
		return aggregate.MultiUnit, nil
	default:
		return aggregate.PropertyType(0), errors.New("unknown property type")
	}
}

// DomainPropertyTypeToProto converts a PropertyType enum from the domain entity to proto.
func DomainPropertyTypeToProto(domainType aggregate.PropertyType) proto.PropertyType {
	switch domainType {
	case aggregate.SingleFamilyHome:
		return proto.PropertyType_SINGLE_FAMILY_HOME
	case aggregate.MultiUnit:
		return proto.PropertyType_MULTI_UNIT
	default:
		return proto.PropertyType_PROPERTY_TYPE_UNSPECIFIED
	}
}

// applyFieldMask applies the fieldMask to the target entity based on the source entity.
func applyFieldMask(target *aggregate.Property, fieldMask *fieldmaskpb.FieldMask, source *proto.RentalProperty) error {
	protoFields := fieldMask.GetPaths()
	for _, path := range protoFields {
		switch path {
		case "property_name":
			target.SetName(source.PropertyName)
		}
	}
	return nil
}
