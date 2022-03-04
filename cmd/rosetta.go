package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/juanvillacortac/rosetta/pkg/generators"
	"github.com/juanvillacortac/rosetta/pkg/model"
	"github.com/yoheimuta/go-protoparser/v4"
)

var (
	proto      = flag.String("proto", "test.proto", "path to the Protocol Buffer file")
	debug      = flag.Bool("debug", false, "debug flag to output more parsing process detail")
	permissive = flag.Bool("permissive", true, "permissive flag to allow the permissive parsing rather than the just documented spec")
	unordered  = flag.Bool("unordered", false, "unordered flag to output another one without interface{}")
)

func run() int {
	reader, err := os.Open(*proto)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open %s, err %v\n", *proto, err)
		return 1
	}
	defer reader.Close()

	got, err := protoparser.Parse(
		reader,
		protoparser.WithDebug(*debug),
		protoparser.WithPermissive(*permissive),
		protoparser.WithFilename(filepath.Base(*proto)),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse, err %v\n", err)
		return 1
	}

	if models, err := model.GetModelsFromPrograms(got); err != nil {
		fmt.Fprintf(os.Stderr, "[Models parsing error]: %v\n", err)
	} else {
		generators.Generate(models)
	}

	// var v interface{}

	// if models, err := model.GetModelsFromPrograms(got); err != nil {
	// 	fmt.Fprintf(os.Stderr, "[Models parsing error]: %v\n", err)
	// } else {
	// 	v = models
	// }

	// v = got
	// if *unordered {
	// 	v, err = protoparser.UnorderedInterpret(got)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "failed to interpret, err %v\n", err)
	// 		return 1
	// 	}
	// }

	// gotJSON, err := json.MarshalIndent(v, "", "  ")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "failed to marshal, err %v\n", err)
	// }
	// fmt.Print(string(gotJSON))
	return 0
}

func Execute() {
	os.Exit(run())
}
