package main

import (
	"os"
	"project-symi-backend/app/infrastructure"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(500)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	infrastructure.Router.Run(":" + port)
}
