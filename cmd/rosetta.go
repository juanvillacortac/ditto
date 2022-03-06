package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/juanvillacortac/rosetta/pkg/program"
)

var (
	source     = flag.String("source", "", "path to the Protocol Buffer file")
	debug      = flag.Bool("debug", false, "debug flag to output more parsing process detail")
	permissive = flag.Bool("permissive", true, "permissive flag to allow the permissive parsing rather than the just documented spec")
)

func run() int {
	flag.Parse()

	if *source == "" {
		fmt.Fprintf(os.Stderr, "You must provide a source \".json\" file\n")
		flag.Usage()
		return 1
	}

	reader, err := os.Open(*source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open %s, err %v\n", *source, err)
		return 1
	}
	defer reader.Close()

	p, err := program.NewProgramFromJson(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}
	if err := p.Parse(
		program.WithDebug(*debug),
		program.WithPermissive(*permissive),
	); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}

	files, err := p.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}
	for _, f := range files {
		path := filepath.Dir(f.Filename)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return 1
		}
		if err := os.WriteFile(f.Filename, []byte(f.Body), os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return 1
		}
	}
	fmt.Println("Done!")
	return 0
}

func Execute() {
	os.Exit(run())
}
