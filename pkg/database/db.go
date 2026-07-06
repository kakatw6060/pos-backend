package database

import (
	"fmt"
	"os"
	"pos-backend/internal/model"

	"gorm.io/driver/postgres"
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
	default:
		return nil, fmt.Errorf("unsupported or missing DB_TYPE. Please set DB_TYPE=postgres")
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
