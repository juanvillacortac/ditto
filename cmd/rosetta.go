package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/juanvillacortac/rosetta/pkg/program"
)

var (
	version = "0.0.0"
	commit  = "XXX"

	source      = flag.String("c", "", "path to the config file (json or yaml)")
	schema      = flag.String("s", "", "path to the schema file (overrides defined in json config file)")
	debug       = flag.Bool("debug", false, "debug flag to output more parsing process detail")
	permissive  = flag.Bool("permissive", true, "permissive flag to allow the permissive parsing rather than the just documented spec")
	showVersion = flag.Bool("v", false, "show version")
	verbose     = flag.Bool("V", false, "Verbose output")
)

func run() int {
	flag.Parse()

	if *showVersion {
		if *verbose {
			fmt.Fprintf(os.Stdout, "Rosetta v%v - Commit hash %v\n", version, commit)
		} else {
			fmt.Fprintf(os.Stdout, "v%v\n", version)
		}
		return 1
	}

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

	p, err := program.NewProgramFromConfigFile(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}

	if *schema != "" {
		p.SchemaFile = *schema
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
