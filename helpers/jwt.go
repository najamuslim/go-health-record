package helpers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Helpers struct{}

func NewHelper() HelperInterface {
	return &Helpers{}
}

// Claims structure to hold JWT claims
type Claims struct {
	jwt.StandardClaims
}

func (h *Helpers) GenerateToken(userID string) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("%v", userID),
		"exp": time.Now().Add(time.Hour * 8).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

// ValidateJWT validates the JWT token
func (h *Helpers) ValidateJWT(tokenString string) (*Claims, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token validation failed: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
