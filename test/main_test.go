package test

import (
	"testing"

	"github.com/certyclick_verify/core"
)

const (
	FilePath     = "../2023-07-13T10_40_18.359468.jpg"
	ExpectedHash = "18ed2f042abf40449b37de7b449424b61866342367fa1a39b5fe0708fc12ccdf"
)

func TestCalculateHash(t *testing.T) {
	// calculate the hash
	hash, err := core.CalculateAllHashes(FilePath)
	if err != nil {
		t.Errorf("Error calculating hash: %s", err)
	}

	// check for each hash if there is a match
	if hash["sha256"] != ExpectedHash {
		t.Errorf("sha256 hash does not match")
	}

	if hash["sha512"] != ExpectedHash {
		t.Errorf("sha512 hash does not match")
	}

	if hash["sha3-256"] != ExpectedHash {
		t.Errorf("sha3-256 hash does not match")
	}

	if hash["sha3-512"] != ExpectedHash {
		t.Errorf("sha3-512 hash does not match")
	}

	if hash["keccak-256"] != ExpectedHash {
		t.Errorf("keccak-256 hash does not match")
	}

	if hash["keccak-512"] != ExpectedHash {
		t.Errorf("keccak-512 hash does not match")
	}

	if hash["sha512/256"] != ExpectedHash {
		t.Errorf("sha512/256 hash does not match")
	}

	if hash["sha512/224"] != ExpectedHash {
		t.Errorf("sha512/224 hash does not match")
	}

	if hash["sha3-224"] != ExpectedHash {
		t.Errorf("sha3-224 hash does not match")
	}

	if hash["sha3-384"] != ExpectedHash {
		t.Errorf("sha3-384 hash does not match")
	}

	if hash["blake2b-256"] != ExpectedHash {
		t.Errorf("blake2b-256 hash does not match")
	}

	if hash["blake2b-512"] != ExpectedHash {
		t.Errorf("blake2b-512 hash does not match")
	}

	if hash["sha256-simd"] != ExpectedHash {
		t.Errorf("sha256-simd hash does not match")
	}

}
