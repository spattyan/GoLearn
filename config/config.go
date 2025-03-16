package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	// Initiliaze Logger
	logger = NewLogger(prefix)
	return logger
}