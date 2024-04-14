package internal

import (
	"flag"
	"fmt"
	"log"

	"github.com/lithammer/dedent"
)

var (
	usage string = fmt.Sprintf(dedent.Dedent(`
		Usage:
		  pg_dump_tree [options]
		
		Options:
		  -h, --help            display this help and exit
		  -f, --file=FILE       %s
		  -o, --output-dir=DIR  %s
	`)[1:], filePathDesc, outputDirDesc)

	outputDir        string
	outputDirDesc    string = fmt.Sprintf("output directory (default: %s)", outputDirDefault)
	outputDirDefault string = "."

	filePath        string
	filePathDesc    string = "input file (default: stdin)"
	filePathDefault string = "-"
)

func flags() {
	flag.StringVar(&filePath, "f", filePathDefault, filePathDesc)
	flag.StringVar(&filePath, "file", filePathDefault, filePathDesc)

	flag.StringVar(&outputDir, "o", outputDirDefault, outputDirDesc)
	flag.StringVar(&outputDir, "output-dir", outputDirDefault, outputDirDesc)

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), usage)
	}

	flag.Parse()
}

func Tree() {
	flags()

	if err := read(filePath, outputDir); err != nil {
		log.Fatal(err.Error())
	}
}
