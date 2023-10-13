package core_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/certyclick_verify/core"
)

func TestCalculateHash(t *testing.T) {
	// Create a temporary file
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal("Error creating temporary file:", err)
	}
	defer os.Remove(file.Name())

	// Write some data to the file
	data := []byte("Hello World!")
	if _, err := file.Write(data); err != nil {
		t.Fatal("Error writing to temporary file:", err)
	}

	// Calculate the hash of the file
	hashed, err := core.CalculateHash(file.Name())
	if err != nil {
		t.Fatal("Error calculating file hash:", err)
	}

	// Compare the calculated hash with the expected hash
	expected := "7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069"
	if fmt.Sprintf("%x", hashed) != expected {
		t.Fatalf("Expected hash %s, got %x", expected, hashed)
	}

	// Close the file
	if err := file.Close(); err != nil {
		t.Fatal("Error closing temporary file:", err)
	}
}
