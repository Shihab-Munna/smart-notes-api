package main

import (
	"log"
	"net/http"
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

	// Define HTTP routes
	http.HandleFunc("/notes", handlers.NoteHandler)
	http.HandleFunc("/notes/", handlers.NoteByIdHandler)

	// Log the successful server startup
	log.Println("Starting server on :5001...")

	// Start the server and check for errors
	if err := http.ListenAndServe(":5001", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Println("Server started successfully on :5001")
	}
}
