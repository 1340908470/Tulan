package handler

import "gorm.io/gorm"

type Handler struct {
	db *gorm.DB
}

var handler Handler

func InitHandler(db *gorm.DB) {
	handler = Handler{db: db}
}

func GetHandler() Handler {
	return handler
}
