package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"smart_note/database"
	"smart_note/handlers"
	"smart_note/middleware"
	"smart_note/models"

	"github.com/gorilla/mux"
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

	// Migrate the schema for both User and Note models
	log.Println("Migrating the database schema...")
	if err := database.DB.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	} else {
		log.Println("Database migration completed.")
	}

	// ROUTER //
	router := mux.NewRouter()

	// Auth route
	router.HandleFunc("/signup", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Protected route
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.JWTAuthMiddleware)
	protected.HandleFunc("/notes", handlers.NoteHandler).Methods("POST", "GET")
	protected.HandleFunc("/notes/{id}", handlers.NoteByIdHandler).Methods("GET", "PUT", "DELETE")

	// END ROUTER LOGIC //

	// Get the port from environment variables (default to 5001 if not set)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5001" // Default port
		log.Printf("PORT environment variable not set, defaulting to port %s", port)
	}

	// Log the successful server startup
	log.Printf("Starting server on :%s...\n", port)

	// Start the server on the specified port
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Printf("Server started successfully on :%s", port)
	}

}
