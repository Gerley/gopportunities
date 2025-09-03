package handler

import (
	"github.com/Gerley/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLooger("handler")
	db = config.GetSQlite()
}
