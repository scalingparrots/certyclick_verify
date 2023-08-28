package core

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	sha256_simd "github.com/minio/sha256-simd"

	black2b_simd "github.com/minio/blake2b-simd"
	"golang.org/x/crypto/sha3"
	"os"
)

// CalculateHash is a utility function to calculate the hash of a file.
// algorithm used: keccak-512
func CalculateHash(filePath string) (string, error) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(fileBytes)
	hashed := hasher.Sum(nil)

	return hex.EncodeToString(hashed), nil
}

// CalculateAllHashes is a utility function to calculate all the hashes of a file.
// algorithms used: sha256, sha512, sha3-256, sha3-512, keccak-256, keccak-512, sha512/256, sha512/224, sha3-224, sha3-384
func CalculateAllHashes(filePath string) (map[string]string, error) {
	// read the file
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// calculate the hashes
	hashes := make(map[string]string)

	// sha256
	hasher := sha256.New()
	hasher.Write(fileBytes)
	hashed := hasher.Sum(nil)
	hashes["sha256"] = hex.EncodeToString(hashed)

	// sha512
	hasher = sha512.New()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["sha512"] = hex.EncodeToString(hashed)

	// sha3-256
	hasher = sha3.New256()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["sha3-256"] = hex.EncodeToString(hashed)

	// sha3-512
	hasher = sha3.New512()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["sha3-512"] = hex.EncodeToString(hashed)

	// keccak-256
	hasher = sha3.NewLegacyKeccak256()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["keccak-256"] = hex.EncodeToString(hashed)

	// keccak-512
	hasher = sha3.NewLegacyKeccak512()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["keccak-512"] = hex.EncodeToString(hashed)

	// sha512/256
	hasher = sha512.New512_256()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["sha512/256"] = hex.EncodeToString(hashed)

	// sha512/224
	hasher = sha512.New512_224()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["sha512/224"] = hex.EncodeToString(hashed)

	// sha3-224
	hasher = sha3.New224()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["sha3-224"] = hex.EncodeToString(hashed)

	// sha3-384
	hasher = sha3.New384()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["sha3-384"] = hex.EncodeToString(hashed)

	// blake2b-256
	hasher = black2b_simd.New256()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["blake2b-256"] = hex.EncodeToString(hashed)

	// blake2b-512
	hasher = black2b_simd.New512()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["blake2b-512"] = hex.EncodeToString(hashed)

	// sha256-simd
	hasher = sha256_simd.New()
	hasher.Write(fileBytes)
	hashed = hasher.Sum(nil)
	hashes["sha256-simd"] = hex.EncodeToString(hashed)

	return hashes, nil
}
