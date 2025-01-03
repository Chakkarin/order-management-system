package main

import (
	"log"
	"net"

	"order-management-system/services/auth-service/internal/config"
	"order-management-system/services/auth-service/internal/controllers"
	"order-management-system/services/auth-service/internal/repositories"
	"order-management-system/services/auth-service/internal/usecases"
	"order-management-system/services/auth-service/proto/auth"

	"google.golang.org/grpc"
)

func main() {

	// Load Configuration
	cfg := config.LoadConfig()

	// Initialize Dependencies
	deps := initDependencies(cfg)

	//
	repo := repositories.NewUserRepository(deps.PgDB)
	authUsecase := usecases.NewUserUsecase(repo, deps)
	handler := controllers.NewUserHandler(authUsecase)

	// Create gRPC Server
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, handler)

	// Start gRPC server
	portgRPC := cfg.GrpcPort
	listener, err := net.Listen("tcp", portgRPC)
	if err != nil {
		log.Fatalf("❌ Failed to listen on port %s: %v", portgRPC, err)
	}
	log.Printf("🥳 Auth-Service gRPC server is running on port %s...", portgRPC)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
