package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		PgDatabase  Database
		Redis       Redis
		RabbitMq    Mq
		ServicePort string
		GrpcPort    string
		ServiceMode string
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}

	Redis struct {
		Host     string
		Port     string
		User     string
		Password string
	}

	Mq struct {
		Host     string
		Port     string
		User     string
		Password string
	}
)

func LoadConfig() *Config {
	if ex := godotenv.Load(".env"); ex != nil {
		log.Panic("Error loading .env file")
	}

	return &Config{
		PgDatabase: Database{
			Host:     getEnv(`PG_AUTH_HOST`),
			Port:     getEnv(`PG_AUTH_PORT`),
			User:     getEnv(`PG_AUTH_USER`),
			Password: getEnv(`PG_AUTH_PASSWORD`),
			DBName:   getEnv(`PG_AUTH_DB_NAME`),
		},
		Redis: Redis{
			Host: getEnv(`REDIS_HOST`),
			Port: getEnv(`REDIS_PORT`),
		},
		RabbitMq: Mq{
			Host:     getEnv(`RABBIT_MQ_HOST`),
			Port:     getEnv(`RABBIT_MQ_PORT`),
			User:     getEnv(`RABBIT_MQ_USER`),
			Password: getEnv(`RABBIT_MQ_PASSWORD`),
		},
		ServicePort: getEnv(`AUTH_SERVICE_PORT`),
		GrpcPort:    getEnv(`GRPC_PORT`),
		ServiceMode: getEnv(`GIN_MODE`),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	log.Panicf(`Error loading .env %v`, key)

	return ""
}
