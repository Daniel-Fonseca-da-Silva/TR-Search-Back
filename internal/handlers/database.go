package handlers

import (
	"fmt"
	"os"

	"tr-search-back/internal/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewUserRepository(db *gorm.DB) *repositories.UserRepository {
	return repositories.NewUserRepository(db)
}
