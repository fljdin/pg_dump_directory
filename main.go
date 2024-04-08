package main

import (
	"os"

	"github.com/fljdin/pg_dump_tree/internal"
)

func main() {
	internal.Tree()
	os.Exit(0)
}
