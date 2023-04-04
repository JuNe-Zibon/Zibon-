package main

import (
	"log"
	"zibon/common/db"
	"zibon/controllers"
	"zibon/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	cors := middlewares.Cors()
	r := gin.Default()
	r.Use(cors)

	DB := db.Connect()
	db.AutoMigrate(DB)

	controllers.RegisterRoutes(r, DB)

	r.Run(":8000")
}
