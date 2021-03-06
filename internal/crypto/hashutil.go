package crypto

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// GenerateHash ...
func GenerateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%X", h.Sum(nil))
}

// HashAndSalt ...
func HashAndSalt(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash)
}

// Compare Passwords
func ComparePasswords(hwpd, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hwpd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}
