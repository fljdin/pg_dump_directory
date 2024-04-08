package internal

import (
	"fmt"
	"os"
)

var (
	input *os.File = os.Stdin
	err   error
)

func read(path string) error {
	if path != "-" {
		input, err = os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer input.Close()
	}

	return nil
}
