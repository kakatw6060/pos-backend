package database

import (
	"fmt"
	"os"
	"pos-backend/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dbType := os.Getenv("DB_TYPE")
	var dialector gorm.Dialector

	switch dbType {
	case "postgres":
		dsn := os.Getenv("DB_URL")
		if dsn == "" {
			return nil, fmt.Errorf("DB_URL environment variable is required for postgres")
		}
		dialector = postgres.Open(dsn)
	case "sqlite":
		fallthrough
	default:
		dbPath := os.Getenv("DB_PATH")
		if dbPath == "" {
			dbPath = "pos.db"
		}
		dialector = sqlite.Open(dbPath)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate models
	err = db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Product{},
		&model.Order{},
		&model.OrderItem{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
