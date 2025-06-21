package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"prosigliere-coding-challenge/config"
	"prosigliere-coding-challenge/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dbConn := config.InitDB()
	config.AutoMigrate(dbConn)

	router := routes.SetupRouter(dbConn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
