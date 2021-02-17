package password

import "golang.org/x/crypto/bcrypt"

// GeneratePasswordHash generates a hash for the password
func GeneratePasswordHash(plainTextPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 14)
	return string(bytes), err
}

// CheckPasswordHash checks the password using the pain text password and comparing it to a previously generated hash 
func CheckPasswordHash(plaintTextPassword, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintTextPassword))
    return err == nil
}