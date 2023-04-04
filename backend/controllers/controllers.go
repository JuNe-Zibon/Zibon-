package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, DB *gorm.DB) {
	H := handler{
		DB: DB,
	}

	auth := r.Group("/auth")
	auth.POST("/register", H.RegisterUser)
}
