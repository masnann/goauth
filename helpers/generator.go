package helpers

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

type Generator struct {
}

type GeneratorInterface interface {
	GenerateHash(password string) (string, error)
	ComparePassword(hash, password string) (bool, error)
}

func NewGenerator() Generator {
	return Generator{}
}

const (
	saltSize    = 16
	keySize     = 32
	timeCost    = 1
	memory      = 64 * 1024
	parallelism = 2
)

func (g Generator) GenerateHash(password string) (string, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, timeCost, memory, parallelism, keySize)
	saltAndHash := append(salt, hash...)
	encodedSaltAndHash := base64.RawStdEncoding.EncodeToString(saltAndHash)

	return encodedSaltAndHash, nil
}

func (g Generator) ComparePassword(hash, password string) (bool, error) {
	decodedSaltAndHash, err := base64.RawStdEncoding.DecodeString(hash)
	if err != nil {
		return false, err
	}

	if len(decodedSaltAndHash) < saltSize {
		return false, errors.New("invalid hash format")
	}

	salt := decodedSaltAndHash[:saltSize]
	existingHash := decodedSaltAndHash[saltSize:]

	computedHash := argon2.IDKey([]byte(password), salt, timeCost, memory, parallelism, keySize)

	if subtle.ConstantTimeCompare(existingHash, computedHash) == 1 {
		return true, nil
	}

	return false, errors.New("password mismatch")
}
