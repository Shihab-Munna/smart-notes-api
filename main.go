package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"smart_note/database"
	"smart_note/handlers"
	"smart_note/models"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Println("Initializing the database connection...")
	database.Init()

	// Check if the DB is properly initialized
	if database.DB == nil {
		log.Fatal("Database connection is nil! Exiting...")
	}

	log.Println("Migrating the Note model...")
	if err := database.DB.AutoMigrate(&models.Note{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	} else {
		log.Println("Migrating Done ....")
	}

	// Get the port from environment variables (default to 5001 if not set)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5001" // Default port
		log.Printf("PORT environment variable not set, defaulting to port %s", port)
	}

	// Define HTTP routes
	http.HandleFunc("/notes", handlers.NoteHandler)
	http.HandleFunc("/notes/", handlers.NoteByIdHandler)

	// Log the successful server startup
	log.Printf("Starting server on :%s...\n", port)

	// Start the server on the specified port
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Printf("Server started successfully on :%s", port)
	}
}
