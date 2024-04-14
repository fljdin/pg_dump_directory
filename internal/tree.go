package internal

import (
	"fmt"
	"os"

	"github.com/fljdin/fragment/languages"
)

var (
	input *os.File = os.Stdin
	err   error
)

func read(path string, dir string) error {
	if path != "-" {
		input, err = os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer input.Close()
	}

	statements := make(chan string)
	go languages.PgSQL.Read(statements, input)

	if dir != "." {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	// TODO: create output files
	entry := dir + "/entrypoint.sql"
	output, err := os.OpenFile(entry, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer output.Close()

	// loop through input statements
	for stmt := range statements {
		statement := Parse(stmt)

		// TODO: redirect statements to the appropriate output file
		if statement.Type == "unknown" {
			output.WriteString(stmt)
		}
	}

	return nil
}
