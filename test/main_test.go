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

}
