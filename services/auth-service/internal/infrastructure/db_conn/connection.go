package dbconn

import (
	"fmt"
	"log"
	"os"
	"services/auth-service/internal/config"
	"services/auth-service/internal/domains/authen/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(conf *config.Database) *gorm.DB {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		conf.Host, conf.User, conf.Password, conf.DBName, conf.Port)

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
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Panicf("❌ Failed to migrate schema: %v", err)
	}

	if ex := tx.Commit().Error; ex != nil {
		tx.Rollback()
		panic(ex)
	}

	return db
}
