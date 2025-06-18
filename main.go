package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"config/db"
	"routes/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	db := db.InitDB()

	db.AutoMigrate(db)

	router := routes.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
