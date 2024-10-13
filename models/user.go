package models

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `json:"password"`
}

// HashPassword hashes the user's password using bcrypt
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	log.Println("Hashed password from DB:", user.Password)
	log.Println("Provided password during login:", providedPassword)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		log.Println("Password comparison failed:", err)
	} else {
		log.Println("Password comparison succeeded")
	}

	return err
}
