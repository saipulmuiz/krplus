package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/pkg/serror"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (cfg *Config) InitPostgres() serror.SError {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname =%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatalf("failed connect to database %+v", err)
		return serror.NewFromError(err)
	}

	err = db.Debug().AutoMigrate(
		models.User{},
		models.CreditLimit{},
		models.Transaction{},
		models.Payment{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database %+v", err)
		return serror.NewFromError(err)
	}

	if db.Migrator().HasTable(&models.User{}) {
		if err := db.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			users := []models.User{
				{FullName: "Budi", LegalName: "Budi", Email: "budi@gmail.com", Password: "budi123", NIK: "2093928392829394"},
				{FullName: "Annisa", LegalName: "Annisa", Email: "annisa@gmail.com", Password: "andi123", NIK: "2093928392829395"},
			}
			if err := db.Create(&users).Error; err != nil {
				log.Printf("Error seeding users: %s", err)
			} else {
				log.Println("Users seeded successfully")
			}
		}
	}

	if db.Migrator().HasTable(&models.CreditLimit{}) {
		if err := db.First(&models.CreditLimit{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			creditLimit := []models.CreditLimit{
				{UserID: 1, Tenor: 1, InitialLimitAmount: 100000, UsedLimitAmount: 0, RemainingLimitAmount: 100000},
				{UserID: 2, Tenor: 1, InitialLimitAmount: 1000000, UsedLimitAmount: 0, RemainingLimitAmount: 1000000},
				{UserID: 1, Tenor: 2, InitialLimitAmount: 200000, UsedLimitAmount: 0, RemainingLimitAmount: 200000},
				{UserID: 2, Tenor: 2, InitialLimitAmount: 1200000, UsedLimitAmount: 0, RemainingLimitAmount: 1200000},
				{UserID: 1, Tenor: 3, InitialLimitAmount: 500000, UsedLimitAmount: 0, RemainingLimitAmount: 500000},
				{UserID: 2, Tenor: 3, InitialLimitAmount: 1500000, UsedLimitAmount: 0, RemainingLimitAmount: 1500000},
				{UserID: 1, Tenor: 6, InitialLimitAmount: 700000, UsedLimitAmount: 0, RemainingLimitAmount: 700000},
				{UserID: 2, Tenor: 6, InitialLimitAmount: 2000000, UsedLimitAmount: 0, RemainingLimitAmount: 2000000},
			}
			if err := db.Create(&creditLimit).Error; err != nil {
				log.Printf("Error seeding creditLimit: %s", err)
			} else {
				log.Println("CreditLimit seeded successfully")
			}
		}
	}

	cfg.DB = db

	// GlobalShutdown.RegisterGracefullyShutdown("database/postgres", func(ctx context.Context) error {
	// 	return cfg.DB.Close()
	// })

	return nil
}
