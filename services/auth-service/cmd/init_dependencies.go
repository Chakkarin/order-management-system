package main

import (
	"order-management-system/services/auth-service/internal/config"
	"order-management-system/services/auth-service/internal/infrastructure"
)

func initDependencies(cfg *config.Config) *config.AppDependencies {
	db := infrastructure.ConnectDB(&cfg.PgDatabase)
	mq := infrastructure.ConnectMQ(&cfg.RabbitMq)
	redis := infrastructure.ConnectRedis(&cfg.Redis)

	return &config.AppDependencies{
		PgDB:     db,
		Redis:    redis,
		RabbitMq: mq,
	}
}
