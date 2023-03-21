package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"example/hello/router"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Server listening on port", port)
	log.Fatal(router.StartServer(port))
}