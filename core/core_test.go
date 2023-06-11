package core_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/certyclick_verify/core"
)

func TestCalculateHash(t *testing.T) {
	// Create a temporary file
	file, err := ioutil.TempFile("", "testfile")
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
	hashed, err := core.CalculateHash(file)
	if err != nil {
		t.Fatal("Error calculating file hash:", err)
	}

	// Compare the calculated hash with the expected hash
	expected := "a69f73cca23a9ac5c8b567dc185a756e97c982164fe25859e0d1dcc1475c80a615b2123af1f5f94c11e3e9402c3ac558f500199d95b6d3e301758586281dcd26"
	if fmt.Sprintf("%x", hashed) != expected {
		t.Fatalf("Expected hash %s, got %x", expected, hashed)
	}

	// Close the file
	if err := file.Close(); err != nil {
		t.Fatal("Error closing temporary file:", err)
	}
}
