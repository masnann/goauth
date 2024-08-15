package helpers

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"go-auth/config"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/argon2"
)

type Generator struct {
}

type GeneratorInterface interface {
	GenerateHash(input string) (string, error)
	CompareHash(hash, input string) (bool, error)
	GenerateJWT(userID int64, email, role string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GenerateRefreshToken(userID int64) (string, error)
	ValidateRefreshToken(tokenString string) (int64, error)
	GenerateOTP(length int) (string, error)
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

func (g Generator) GenerateHash(input string) (string, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(input), salt, timeCost, memory, parallelism, keySize)
	saltAndHash := append(salt, hash...)
	encodedSaltAndHash := base64.RawStdEncoding.EncodeToString(saltAndHash)

	return encodedSaltAndHash, nil
}

func (g Generator) CompareHash(hash, input string) (bool, error) {
	decodedSaltAndHash, err := base64.RawStdEncoding.DecodeString(hash)
	if err != nil {
		return false, err
	}

	if len(decodedSaltAndHash) < saltSize {
		return false, errors.New("invalid hash format")
	}

	salt := decodedSaltAndHash[:saltSize]
	existingHash := decodedSaltAndHash[saltSize:]

	computedHash := argon2.IDKey([]byte(input), salt, timeCost, memory, parallelism, keySize)

	if subtle.ConstantTimeCompare(existingHash, computedHash) == 1 {
		return true, nil
	}

	return false, errors.New("input mismatch")
}

func (g Generator) GenerateJWT(userID int64, email, role string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"role":   role,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (g Generator) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (g Generator) GenerateRefreshToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (g Generator) ValidateRefreshToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userID"].(float64)
		if !ok {
			return 0, errors.New("invalid user_id type")
		}
		return int64(userID), nil
	} else {
		return 0, errors.New("invalid refresh token")
	}
}

func (g Generator) GenerateOTP(length int) (string, error) {
	otp := make([]byte, length)
	_, err := rand.Read(otp)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		otp[i] = (otp[i] % 10) + '0'
	}

	return string(otp), nil
}

func (g Generator) CompareOTP(otpHash, otp string) (bool, error) {
	decodedSaltAndHash, err := base64.RawStdEncoding.DecodeString(otpHash)
	if err != nil {
		return false, err
	}

	if len(decodedSaltAndHash) < saltSize {
		return false, errors.New("invalid hash format")
	}

	salt := decodedSaltAndHash[:saltSize]
	existingHash := decodedSaltAndHash[saltSize:]

	computedHash := argon2.IDKey([]byte(otp), salt, timeCost, memory, parallelism, keySize)

	if subtle.ConstantTimeCompare(existingHash, computedHash) == 1 {
		return true, nil
	}

	return false, errors.New("otp mismatch")
}
