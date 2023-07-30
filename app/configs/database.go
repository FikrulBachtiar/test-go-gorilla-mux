package configs

import (
	"fmt"
	"log"
	"test-golang/app/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *EnvironmentVariable) *gorm.DB {
	dataSource := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=%s`, config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_SSL_MODE)
	dbGorm, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("DB GORM Error Open => %s", err))
	}

	dbGorm.AutoMigrate(&models.Customer{})
	dbGorm.AutoMigrate(&models.Family{})

	db, err := dbGorm.DB()
	if err != nil {
		log.Fatal(fmt.Sprintf("DB Error Open => %s", err))
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(20 * time.Minute)

	if err = db.Ping(); err != nil {
		log.Fatal(fmt.Sprintf("Connection DB Error => %s", err))
	}

	return dbGorm
}