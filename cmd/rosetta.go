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

	config      = flag.String("c", "config.yml", "path to the config file (json or yaml)")
	schema      = flag.String("s", "schema.yml", "path to the schema file (overrides defined in json config file)")
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

	if *config == "" {
		fmt.Fprintf(os.Stderr, "You must provide a config file\n")
		flag.Usage()
		return 1
	}

	reader, err := os.Open(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open %s, err %v\n", *config, err)
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

	if err := p.Parse(); err != nil {
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
