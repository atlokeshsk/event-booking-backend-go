package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plain text password and returns its bcrypt hashed version.
// It uses a cost factor of 14 for the bcrypt algorithm, which determines the computational complexity.
// The function returns the hashed password as a string and an error if the hashing process fails.
func HashPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(encryptedPassword), err
}

// CheckHashPassword compares a hashed password with a plain text password
// to determine if they match.
//
// Parameters:
// - hashedPassword: The hashed password string.
// - password: The plain text password string.
//
// Returns:
// - bool: True if the hashed password matches the plain text password, false otherwise.
func CheckHashPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
