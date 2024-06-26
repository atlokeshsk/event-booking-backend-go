package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const signedKey = "amarkalam"

// GenerateJwtToken generates a JWT token for a given email and user ID.
// The token is signed using the HS256 signing method and includes claims for the email, user ID, and expiration time.
//
// Parameters:
//
//	email (string): The email address to include in the token claims.
//	userID (int64): The user ID to include in the token claims.
//
// Returns:
//
//	string: The signed JWT token as a string.
//	error: An error if there was an issue signing the token.
func GenerateJwtToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": userID,
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(signedKey))
}

// VerifyToken parses and validates a JWT token, returning the user ID if successful.
// It returns an error if the token is invalid or cannot be parsed.
//
// Parameters:
//   - token: A string representing the JWT token to be verified.
//
// Returns:
//   - int64: The user ID extracted from the token claims if the token is valid.
//   - error: An error if the token is invalid or cannot be parsed.
func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
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
