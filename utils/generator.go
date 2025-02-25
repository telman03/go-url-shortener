package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	shortCode := make([]byte, length)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortCode)
}