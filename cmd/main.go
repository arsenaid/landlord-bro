package main

import (
	"context"
	"landlord-bro/infrastructure/database"
	grpcserver "landlord-bro/infrastructure/grpc"
	proto "landlord-bro/interfaces/api/gen/v1"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
)

func main() {
	boot := rkboot.NewBoot()

	// Try logger
	rkBootLogger := rkentry.GlobalAppCtx.GetLoggerEntry("logger")
	rkBootConfig := rkentry.GlobalAppCtx.GetConfigEntry("db-config")

	dbConfig := database.Config{
		Dialect:      rkBootConfig.GetString("dialect"),
		Host:         rkBootConfig.GetString("host"),
		Port:         rkBootConfig.GetInt("port"),
		Username:     rkBootConfig.GetString("user"),
		Password:     rkBootConfig.GetString("password"),
		DatabaseName: rkBootConfig.GetString("dbname"),
		SSLMode:      rkBootConfig.GetString("sslMode"),
		OtherOptions: nil,
	}

	dbClient, err := database.NewDBClient(dbConfig)
	if err != nil {
		panic("failed to connect to database")
	}

	// register grpc
	entry := rkgrpc.GetGrpcEntry("landlord-pro-grpc")
	grpcServer := &grpcserver.Server{DB: dbClient, Logger: rkBootLogger}
	entry.AddRegFuncGrpc(
		grpcServer.RegisterRentalPropertyServer,
		// grpcServer.RegisterTenantServer,
		// grpcServer.RegisterLeaseServer,
	)
	entry.AddRegFuncGw(
		proto.RegisterPropertyServiceHandlerFromEndpoint,
		// proto.RegisterTenantServiceHandlerFromEndpoint,
		// proto.RegisterLeaseServiceHandlerFromEndpoint,
	)
	entry.AddUnaryInterceptors(
		grpcmiddleware.ChainUnaryServer(
			grpcvalidator.UnaryServerInterceptor(),
		),
	)
	// Bootstrap
	boot.Bootstrap(context.Background())

	// Wait for shutdown signal
	rkentry.GlobalAppCtx.WaitForShutdownSig()

	// Interrupt gin entry
	entry.Interrupt(context.Background())
}
