package jwtutils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetSecretKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "default-secret-key"
	}
	return []byte(key)
}

type JWTClaim struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Nik string `json:"nik"`
	Role  string `json:"role"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func CreateAccessToken(ID uint, Name string, Nik string, Role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &JWTClaim{
		ID:     ID,
		Name: Name,
		Nik: 	Nik,
		Role:   Role,
		Type: 	"access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetSecretKey())
}

func CreateRefreshToken(ID uint) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)

	claims := &JWTClaim{
		ID:   ID,
		Type: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetSecretKey())
}

func VerifyToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return GetSecretKey(), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil
}

func GetCurrentUser(c *gin.Context) (*JWTClaim, bool) {
	userValue, exists := c.Get("user")
	if !exists {
		return nil, false
	}

	user, ok := userValue.(*JWTClaim)
	if !ok {
		return nil, false
	}

	return user, true

}

func GetCurrentUserID(c *gin.Context) (uint, bool) {
	user, ok := GetCurrentUser(c)
	if !ok {
		return 0, false
	}

	return user.ID, true
}