package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize SQLite
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}
	return nil
}

func GetSQlite() *gorm.DB {
	return db
}

func GetLooger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
