package config

import (
	"fmt"
	"os"
)

func LoadDataSourceName() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// cfg := dbDriver{
	// 	Username: os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Host: os.Getenv("DB_HOST"),
	// 	Port: os.Getenv("DB_PORT"),
	// 	DBName: os.Getenv("DB_NAME"),
	// }

	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)

	// // db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}