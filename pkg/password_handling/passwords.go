package password_handling

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/big"
	"strings"
)

func HashAndSaltPassword(password string) string {
	h, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	s := string(h)
	return s
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func IsPasswordValid(password string, minLength int, complexityNeeded bool) bool {
	if len(password) < minLength {
		return false
	}
	if complexityNeeded {
		// Check for complexity
		// 1. At least one uppercase letter
		if strings.ToUpper(password) == password {
			return false
		}
		// 2. At least one lowercase letter
		if strings.ToLower(password) == password {
			return false
		}
		// 3. At least one number
		if strings.IndexAny(password, "0123456789") == -1 {
			return false
		}
		// 4. At least one special character
		if strings.IndexAny(password, "!@#$%^&*") == -1 {
			return false
		}
	}
	return true
}

func GenerateRandomPassword(length int) string {
	// Define the character set to use for generating the random password. Avoid characters that can be used
	// to escape strings in a shell or SQL injection attacks or create other security vulnerabilities such as `, ", \, and ;
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	charSetLen := int64(len(charset))
	randomBytes := make([]byte, length)
	for i := range randomBytes {
		randomIndex := generateRandomNumber(charSetLen)
		randomBytes[i] = charset[randomIndex]
	}
	return string(randomBytes)
}

func generateRandomNumber(max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return 0
	}
	return n.Int64()
}
