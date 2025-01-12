package server

import (
	"services/auth-service/internal/domains/authen/controllers"
	"services/auth-service/internal/domains/authen/repositories"
	"services/auth-service/internal/domains/authen/usecases"
	mqconn "services/auth-service/internal/infrastructure/mq_conn"
	"services/auth-service/shared/proto/authen"
)

func (s *ginServer) initAuthenRouter() {

	// สร้าง Repository
	userRepo := repositories.NewUserRepository(s.dbconn)
	// เรียกใช้ util ของ rabbitmq
	mqUtils := mqconn.NewMqUtils(s.rabbitMqConn)
	// สร้าง Usecase
	authUsecase := usecases.NewUserUsecase(&usecases.AuthenUsecaseDependencies{
		Repo:     userRepo,
		Redis:    s.redisConn,
		RabbitMq: mqUtils,
	})

	// gRPC
	grpcHandler := controllers.NewAuthenGRPCHandler(authUsecase)
	authen.RegisterAuthServiceServer(s.grpcServer, grpcHandler)

	// REST
	httpHandler := controllers.NewAuthenHTTPHandler(authUsecase)
	authGroup := s.app.Group("/v1/auth")
	{
		authGroup.POST("/register", httpHandler.Register)
	}
}
