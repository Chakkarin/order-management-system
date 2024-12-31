package main

import (
	"log"
	"net"
	"os"

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
		log.Panic("❌ No .env file found, loading from system env")
	}

	log.Println("✅ loaded .env ...")
}

func main() {

	mq := infrastructure.ConnectMQ()

	// Connect to redis
	redis := infrastructure.ConnectRedis()

	// Connect to database
	db := infrastructure.ConnectDB()

	// Migrate schema
	// db.AutoMigrate(&domain.User{})

	// Initialize Dependencies
	userRepo := repositories.NewUserRepository(db)
	authUsecase := usecases.NewUserUsecase(userRepo, redis, mq)
	authController := controllers.NewUserHandler(authUsecase)

	// Create gRPC Server
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authController)

	// Start gRPC server
	portgRPC := os.Getenv("PORT_GRPC")
	listener, err := net.Listen("tcp", portgRPC)
	if err != nil {
		log.Fatalf("❌ Failed to listen on port %s: %v", portgRPC, err)
	}
	log.Printf("🥳 Auth-Service gRPC server is running on port %s...", portgRPC)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
