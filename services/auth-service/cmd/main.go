package main

import (
	"log"
	"net"

	"order-management-system/services/auth-service/internal/controllers"
	"order-management-system/services/auth-service/internal/infrastructure"
	"order-management-system/services/auth-service/internal/repositories"
	"order-management-system/services/auth-service/internal/usecases"
	"order-management-system/services/auth-service/proto/auth"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading from system env")
	}

	log.Println("loaded .env ...")
}

func main() {
	// Connect to database
	db := infrastructure.ConnectDB()

	// Connect to redis
	redis := infrastructure.ConnectRedis()

	// Migrate schema
	// db.AutoMigrate(&domain.User{})

	// Initialize Dependencies
	userRepo := repositories.NewUserRepository(db)
	authUsecase := usecases.NewUserUsecase(userRepo, redis)
	authController := controllers.NewUserHandler(authUsecase)

	// Create gRPC Server
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authController)

	// Start gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	log.Println("Auth-Service gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
