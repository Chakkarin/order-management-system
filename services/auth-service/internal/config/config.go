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
		log.Fatal("Error loading .env file")
	}

	return &Config{
		PgDatabase: Database{
			Host:     os.Getenv(`PG_AUTH_HOST`),
			Port:     os.Getenv(`PG_AUTH_PORT`),
			User:     os.Getenv(`PG_AUTH_USER`),
			Password: os.Getenv(`PG_AUTH_PASSWORD`),
			DBName:   os.Getenv(`PG_AUTH_DB_NAME`),
		},
		Redis: Redis{
			Host: os.Getenv(`REDIS_HOST`),
			Port: os.Getenv(`REDIS_PORT`),
		},
		RabbitMq: Mq{
			Host:     os.Getenv(`RABBIT_MQ_HOST`),
			Port:     os.Getenv(`RABBIT_MQ_USER`),
			User:     os.Getenv(`RABBIT_MQ_PASSWORD`),
			Password: os.Getenv(`RABBIT_MQ_PORT`),
		},
		ServicePort: os.Getenv(`AUTH_SERVICE_PORT`),
	}
}
