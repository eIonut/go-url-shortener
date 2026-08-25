package helpers

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomCode() (string, error) {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 7)

	for i := range result {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", err
		}
		result[i] = alphabet[randomIndex.Int64()]
	}

	return string(result), nil
}
