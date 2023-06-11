package core

import (
	"os"
)

// OpenFile is a utility function to open a file.
func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
