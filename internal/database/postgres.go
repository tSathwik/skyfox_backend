package database 

import (
	"fmt"
	"log"
	"skyfox_backend/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler interface {
	NewDatabase() *gorm.DB
}

type dbHandler struct {
	config config.DBConfig
}

func NewDBHandler(cfg *config.Config) *dbHandler {
	return &dbHandler{config:cfg.Database}
}

func (dh *dbHandler) NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",dh.config.Host, dh.config.User, dh.config.Password, dh.config.Name, dh.config.Port)
	log.Printf("Connecting to database with DSN: %s", dsn)
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Database connected successfully")
	return db
}