package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const signedKey = "amarkalam"

func GenerateJwtToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": userID,
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(signedKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected siginin method")
		}
		return []byte(signedKey), nil
	})
	if err != nil {
		return 0, errors.New("unable to parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	//	email := claims["email"].(string)
	userID := int64(claims["user_id"].(float64))
	return userID, nil
}
