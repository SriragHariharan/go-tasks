package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var jwtSecret = []byte("your-secret-key") // move to env in real apps

// custom claims
type Claims struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Generate token
func GenerateJwt(userId bson.ObjectID, userEmail string) (string, error) {
	claims := Claims{
		UserID: userId.Hex(),
		Email:  userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gotasks",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

// Verify token
func VerifyJwt(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// ensure correct signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.UserID, nil
}