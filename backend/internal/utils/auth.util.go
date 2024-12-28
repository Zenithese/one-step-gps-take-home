package utils

import (
	"errors"
	"os"
	"time"
	"strconv"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type TokenClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID, username string) (string, error) {
	claims := TokenClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GetUserIDFromRequest(r *http.Request) (int, error) {
    cookie, err := r.Cookie("auth_token")
    if err != nil {
        return 0, errors.New("no auth token found")
    }

    claims, err := ValidateJWT(cookie.Value)
    if err != nil {
        return 0, errors.New("invalid token")
    }

    userID, err := strconv.Atoi(claims.UserID)
    if err != nil {
        return 0, errors.New("invalid user ID in token")
    }

    return userID, nil
}