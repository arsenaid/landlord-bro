package grpc

import (
	"landlord-bro/domain/service"
	"landlord-bro/infrastructure/database/repository"
	"landlord-bro/infrastructure/grpc/handler"
	proto "landlord-bro/interfaces/api/gen/v1"

	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Logger *rkentry.LoggerEntry
}

func (s *Server) RegisterRentalPropertyServer(server *grpc.Server) {
	propertyRepo := repository.NewPropertyRepository(s.DB)
	propertyService := service.NewPropertyService(propertyRepo)
	grpcHandler := handler.NewPropertyGrpcHandler(propertyService, s.Logger)

	proto.RegisterPropertyServiceServer(server, grpcHandler)
}

// func (s *Server) RegisterTenantServer(server *grpc.Server) {
// 	tenantRepo := repository.NewTenantRepository(s.DB)
// 	tenantService := service.NewTenantService(tenantRepo)
// 	grpcHandler := handler.NewTenantGrpcHandler(tenantService, s.Logger)

// 	proto.RegisterTenantServiceServer(server, grpcHandler)
// }

// func (s *Server) RegisterLeaseServer(server *grpc.Server) {
// 	leaseRepo := repository.NewLeaseRepository(s.DB)
// 	propertyRepo := repository.NewPropertyRepository(s.DB)
// 	leaseService := service.NewLeaseService(leaseRepo, propertyRepo)
// 	grpcHandler := handler.NewLeaseGrpcHandler(leaseService, s.Logger)

// 	proto.RegisterLeaseServiceServer(server, grpcHandler)
// }
