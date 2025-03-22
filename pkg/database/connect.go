package database

import (
	"fmt"
	"log"

	"github.com/The-EpaG/Go-BasicServer/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func CreateConnection() (*gorm.DB, error) {
	cfg := config.Get()

	var dsn string
	var dialect gorm.Dialector

	switch cfg.DB.Driver {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Rome",
			cfg.DB.Host,
			cfg.DB.Username,
			cfg.DB.Password,
			cfg.DB.Name,
			cfg.DB.Port,
		)
		dialect = postgres.Open(dsn)
	case "sqlite":
		dsn = cfg.DB.Path
		dialect = sqlite.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database type %s", cfg.DB.Driver)
	}
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %v", err)
	}
	return db, nil
}
