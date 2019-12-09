package main

import (
	"os"
	"project-symi-backend/app/infrastructure"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	infrastructure.Router.Run(":" + port)
}
