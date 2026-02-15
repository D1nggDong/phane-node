package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

// Config for Argon2 hashing
const (
	saltLen = 16
	time    = 1
	memory  = 64 * 1024
	threads = 4
	keyLen  = 32
)

// Register creates a hashed password to store in the vault
func Register(password string) (string, error) {
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	// Format: salt.hash
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	
	return fmt.Sprintf("%s.%s", b64Salt, b64Hash), nil
}

// Login checks the provided password against the stored hash
func Login(password, storedHash string) (bool, error) {
	parts := strings.Split(storedHash, ".")
	if len(parts) != 2 {
		return false, errors.New("invalid hash format")
	}

	salt, _ := base64.RawStdEncoding.DecodeString(parts[0])
	originalHash, _ := base64.RawStdEncoding.DecodeString(parts[1])

	comparisonHash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	if base64.RawStdEncoding.EncodeToString(comparisonHash) == base64.RawStdEncoding.EncodeToString(originalHash) {
		return true, nil
	}

	return false, nil
}
