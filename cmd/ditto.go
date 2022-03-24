package cmd

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/juanvillacortac/ditto/pkg/program"
)

var (
	version = "0.0.0"
	commit  = "XXX"

	config      = flag.String("c", "config.yml", "Path to the config file (json or yaml)")
	output      = flag.String("o", "", "Path to the output base path")
	schema      = flag.String("s", "", "Path to the schema file (overrides defined in json config file) (default \"schema.yml\" if not defined in config)")
	rm          = flag.Bool("rm", false, "Clean output path before generating")
	showVersion = flag.Bool("v", false, "Show version")
	verbose     = flag.Bool("V", false, "Verbose output")
)

func run() int {
	flag.Parse()

	if *showVersion {
		if *verbose {
			fmt.Fprintf(os.Stdout, "Ditto v%v - Commit hash %v\n", version, commit)
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
	} else if p.SchemaFile == "" {
		p.SchemaFile = "schema.yml"
	}

	if *output != "" {
		p.OutputBasePath = *output
	}

	if err := p.Parse(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}

	basePath, _ := filepath.Abs(path.Dir(*config))
	files, err := p.Generate(basePath, *verbose)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}

	if *rm && p.OutputBasePath != "" {
		if err := os.RemoveAll(p.OutputBasePath); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return 1
		}
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
