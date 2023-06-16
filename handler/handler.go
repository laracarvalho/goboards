package handler

import (
	"github.com/laracarvalho/goboards/config"
	"gorm.io/gorm"
)


var (
	logger *config.Logger
	db *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
}