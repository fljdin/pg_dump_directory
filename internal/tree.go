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

func read(path string) error {
	if path != "-" {
		input, err = os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer input.Close()
	}

	statements := make(chan string)
	go languages.PgSQL.Read(statements, input)

	for stmt := range statements {
		statement := Parse(stmt)

		// TODO: redirect statements to the appropriate output file
		if statement.Name == "" {
			fmt.Print(statement.Raw)
		}
	}

	return nil
}
