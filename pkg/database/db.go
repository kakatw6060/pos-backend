package database

import (
	"fmt"
	"os"
	"strings"
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
		// Force simple protocol to avoid prepared statement errors with Supabase Pooler
		if !strings.Contains(dsn, "prefer_simple_protocol=true") {
			separator := "?"
			if strings.Contains(dsn, "?") {
				separator = "&"
			}
			dsn = fmt.Sprintf("%s%sprefer_simple_protocol=true", dsn, separator)
		}
		dialector = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported or missing DB_TYPE. Please set DB_TYPE=postgres")
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: false,
	})
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
