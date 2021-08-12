package controllers

import "gorm.io/gorm"

type Controller struct {
	db *gorm.DB
}

// Create new controller
func NewController(db *gorm.DB) *Controller {
	return &Controller{db}
}

// Message example
type Message struct {
	Error   int    `json:"error_status" example:"400"`
	Message string `json:"message" example:"message"`
}
