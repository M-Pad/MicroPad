package core

import (
	"os"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func WriteFile(path string, bytes []byte) error {
	return os.WriteFile(path, bytes, 0644)
}
