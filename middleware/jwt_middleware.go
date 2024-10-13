package middleware

import (
	"log"
	"net/http"
	"os"
	"smart_note/utils"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// Secret key for JWT validation
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// JWTAuthMiddleware validates the JWT in the request
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "Authorization header missing")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			log.Printf("JWT token is invalid: %v", err)
			utils.WriteJSONError(w, http.StatusForbidden, "Invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
