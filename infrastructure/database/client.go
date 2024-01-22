package database

import (
	"fmt"
	"landlord-bro/infrastructure/database/model"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Dialect      string
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
	SSLMode      string
	OtherOptions map[string]string
}

func NewDBClient(config Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch config.Dialect {
	case "sqlite":
		db, err = gorm.Open(
			sqlite.Open(config.DatabaseName), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			},
		)
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			config.Host, config.Username, config.Password, config.DatabaseName, config.Port, config.SSLMode,
		)

		for key, value := range config.OtherOptions {
			dsn += " " + key + "=" + value
		}

		db, err = gorm.Open(
			postgres.Open(dsn), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			},
		)
	default:
		return nil, fmt.Errorf("unsupported database dialect: %s", config.Dialect)
	}

	if err != nil {
		return nil, err
	}

	err = initSchema(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// InitSchema will create & migrate the tables, then make some relationships if necessary
func initSchema(db *gorm.DB) error {
	// Initialize the PropertyRepository
	// err := db.AutoMigrate(&model.Property{}, &model.Tenant{}, &model.Lease{})
	err := db.AutoMigrate(&model.Property{}, &model.PropertyUnit{})
	if err != nil {
		return errors.Wrap(err, "failed to create database.")
	}

	return nil
}
