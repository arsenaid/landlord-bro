package handler

import (
	"context"
	"landlord-bro/domain/service"
	"landlord-bro/infrastructure/grpc/translator"

	proto "landlord-bro/interfaces/api/gen/v1"

	"github.com/google/uuid"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkgrpcerr "github.com/rookie-ninja/rk-grpc/v2/boot/error"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PropertyGrpcHandler struct {
	propertyService *service.PropertyService
	logger          *rkentry.LoggerEntry
}

func NewPropertyGrpcHandler(propertyService *service.PropertyService, logger *rkentry.LoggerEntry) *PropertyGrpcHandler {
	return &PropertyGrpcHandler{propertyService: propertyService, logger: logger}
}

func (s *PropertyGrpcHandler) CreateProperty(ctx context.Context, req *proto.CreatePropertyRequest) (*proto.RentalProperty, error) {
	property, err := translator.FromProtoToEntity(req.GetProperty())
	if err != nil {
		return nil, rkgrpcerr.InvalidArgument("failed to parse request.", err).Err()
	}

	err = s.propertyService.CreateProperty(ctx, property)
	if err != nil {
		return nil, rkgrpcerr.Internal("failed to create property.", err).Err()
	}

	return translator.ToProtoFromEntity(property), nil
}

func (s *PropertyGrpcHandler) UpdateProperty(ctx context.Context, req *proto.UpdatePropertyRequest) (*proto.RentalProperty, error) {
	property, err := translator.FromProtoToEntity(req.GetProperty())
	if err != nil {
		return nil, rkgrpcerr.InvalidArgument("failed to parse request.", err).Err()
	}

	err = s.propertyService.UpdateProperty(ctx, property)
	if err != nil {
		return nil, rkgrpcerr.Internal("failed to update property.", err).Err()
	}

	return translator.ToProtoFromEntity(property), nil
}

func (s *PropertyGrpcHandler) DeleteProperty(ctx context.Context, req *proto.DeletePropertyRequest) (*emptypb.Empty, error) {
	err := s.propertyService.DeleteProperty(ctx, uuid.MustParse(req.PropertyId))

	if err != nil {
		return nil, rkgrpcerr.Internal("failed to delete property.", err).Err()
	}

	return &emptypb.Empty{}, nil
}

func (s *PropertyGrpcHandler) GetProperty(ctx context.Context, req *proto.GetPropertyRequest) (*proto.RentalProperty, error) {
	property, err := s.propertyService.GetPropertyByID(ctx, uuid.MustParse(req.PropertyId))

	if err == service.ErrPropertyNotFound {
		return nil, rkgrpcerr.NotFound("property not found.", err).Err()
	}

	if err != nil {
		return nil, rkgrpcerr.Internal("failed to fetch property.", err).Err()
	}

	return translator.ToProtoFromEntity(property), nil
}

func (s *PropertyGrpcHandler) ListProperties(ctx context.Context, req *proto.ListPropertiesRequest) (*proto.ListPropertiesResponse, error) {
	var protoProperties []*proto.RentalProperty

	properties, nextPageToken, err := s.propertyService.ListProperties(ctx, req.PageToken, req.PageSize)

	if err != nil {
		return nil, rkgrpcerr.Internal("failed to list properties.", err).Err()
	}

	for _, property := range properties {
		protoProperties = append(protoProperties, translator.ToProtoFromEntity(property))
	}

	return &proto.ListPropertiesResponse{Properties: protoProperties, NextPageToken: nextPageToken}, nil
}
