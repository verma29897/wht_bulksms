package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/verma29897/bulksms/db"
	"github.com/verma29897/bulksms/models"
	routes "github.com/verma29897/bulksms/routers"
)

func main() {
	_ = godotenv.Load()
	r := gin.Default()
	r.Use(cors.Default())
	// Initialize GORM and auto-migrate core models; exit on failure if DB_URL is set
	if os.Getenv("DB_URL") != "" {
		if err := db.InitGorm(&models.User{}, &models.Account{}); err != nil {
			log.Fatal("GORM init failed:", err)
		}
	}
	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
