package main

// import (
// 	"services/auth-service/internal/config"
// 	"services/auth-service/internal/infrastructure"
// 	"services/auth-service/internal/infrastructure/mq"
// )

// func initDependencies(cfg *config.Config) *config.AppDependencies {
// 	db := infrastructure.ConnectDB(&cfg.PgDatabase)
// 	mq := mq.ConnectMQ(&cfg.RabbitMq)
// 	redis := infrastructure.ConnectRedis(&cfg.Redis)

// 	return &config.AppDependencies{
// 		PgDB:     db,
// 		Redis:    redis,
// 		RabbitMq: mq,
// 	}
// }
