package main

import (
	"log"
	"net"
	"services/auth-service/internal/config"
	"services/auth-service/internal/domains/authen/controllers"
	"services/auth-service/internal/domains/authen/repositories"
	"services/auth-service/internal/domains/authen/usecases"
	"services/auth-service/internal/infrastructure"
	"services/auth-service/internal/infrastructure/mq"
	"services/auth-service/internal/proto/authen"

	"google.golang.org/grpc"
)

func main() {

	// Load Configuration
	cfg := config.LoadConfig()

	// Initialize Dependencies
	// deps := initDependencies(cfg)

	db := infrastructure.ConnectDB(&cfg.PgDatabase)
	redis := infrastructure.ConnectRedis(&cfg.Redis)
	mqConn := mq.ConnectMQ(&cfg.RabbitMq)
	mqx := mq.NewMQReceiver(mqConn)

	//
	repo := repositories.NewAuthenRepository(db)
	authUsecase := usecases.NewUserUsecase(&usecases.AuthenUsecaseDependencies{
		Repo:     repo,
		Redis:    redis,
		RabbitMq: mqx,
	})
	handler := controllers.NewUserHandler(authUsecase)

	// Create gRPC Server
	grpcServer := grpc.NewServer()
	authen.RegisterAuthServiceServer(grpcServer, handler)

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
