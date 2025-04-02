package main

import (
	"fmt"
	"os"

	parse_template "github.com/cocreators-ee/parse-template"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Printf("  %s <source file> [--name=value]\n", os.Args[0])
	fmt.Println("")
	fmt.Printf("More information at https://github.com/cocreators-ee/parse-template\n")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	source, td := parse_template.GetTemplateData(os.Args, os.Environ())
	contents, err := os.ReadFile(source)
	if err != nil {
		panic(err)
	}

	parse_template.CompileTemplate(string(contents[:]), td, os.Stdout)
}
