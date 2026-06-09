package main

import (
	"os"
	"path/filepath"
)

func main() {
	panic("A problem")

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)

	if err != nil {
		panic(err)
	}
}
