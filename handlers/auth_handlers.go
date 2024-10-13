package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"smart_note/database"
	"smart_note/models"
	"smart_note/utils"
	"time"

	"github.com/golang-jwt/jwt"
)

// Initialize JWT secret from environment variables
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the JSON request body into the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Hash the user password
	if err := user.HashPassword(user.Password); err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Error hashing password")
		return
	}

	// Log the hashed password for debugging
	log.Println("Hashed password after bcrypt:", user.Password)

	// Create the user in the database
	if err := database.DB.Create(&user).Error; err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Error creating user or user already exists")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode the incoming JSON payload
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var user models.User
	// Find user by email in the database
	if err := database.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		log.Println(err)
		utils.WriteJSONError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Verify the provided password with the hashed password
	if err := user.CheckPassword(credentials.Password); err != nil {
		log.Println("Hashed password from DB:", user.Password)
		log.Println("Provided plain-text password:", credentials.Password)
		utils.WriteJSONError(w, http.StatusUnauthorized, "Invalid email or password II")
		return
	}

	// Generate a JWT token for the authenticated user
	token, err := generateJWT(user)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

// generateJWT generates a JWT token for the given user
func generateJWT(user models.User) (string, error) {
	// Set token expiration time from environment or default to 20 min.
	expirationTime := time.Minute * 20 
	if expStr := os.Getenv("JWT_EXPIRATION"); expStr != "" {
		if expDuration, err := time.ParseDuration(expStr); err == nil {
			expirationTime = expDuration
		}
	}

	// Create the JWT token with claims (user ID and expiration time)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(expirationTime).Unix(),
	})

	// Sign the token using the secret
	return token.SignedString(jwtSecret)
}
