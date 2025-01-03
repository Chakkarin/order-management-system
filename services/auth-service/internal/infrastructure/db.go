package infrastructure

import (
	"log"
	"order-management-system/services/auth-service/internal/config"
	"order-management-system/services/auth-service/internal/domain"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(conf *config.Database) *gorm.DB {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Panic("❌ DATABASE_URL is not set")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:        time.Second, // Slow SQL threshold
			LogLevel:             logger.Info, // Log level
			ParameterizedQueries: true,        // Don't include params in the SQL log
			Colorful:             true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panicf("❌ Failed to connect to database: %v", err)
	}

	log.Println("✅ connected to database...")

	tx := db.Begin()

	// set database auto uuid
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`); err.Error != nil {
		log.Panicf("❌ Failed to create extension uuid-ossp: %v", err.Error)
	}

	// Migrate schema
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Panicf("❌ Failed to migrate schema: %v", err)
	}

	if ex := tx.Commit().Error; ex != nil {
		tx.Rollback()
		panic(ex)
	}

	return db
}
