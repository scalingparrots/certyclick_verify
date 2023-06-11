package core

import (
	"io"
	"os"

	"golang.org/x/crypto/sha3"
)

// CalculateHash is a utility function to calculate the hash of a file.
func CalculateHash(file *os.File) ([]byte, error) {
	hasher := sha3.New512()
	if _, err := io.Copy(hasher, file); err != nil {
		return nil, err
	}
	hashed := hasher.Sum(nil)
	return hashed, nil
}
