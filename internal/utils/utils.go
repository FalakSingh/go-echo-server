package utils

import (
	"echo-server/config"
	"maps"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func ComparePassword(hashedPassword, givenPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(givenPassword))
	return err == nil
}

func CreateJWT(payload map[string]any) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	maps.Copy(claims, payload)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Env.JWT_SECRET))		
}

