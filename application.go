package main

import (
	"os"
	"project-symi-backend/app/infrastructure"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	print(err)
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	infrastructure.Router.Run(port)
}
