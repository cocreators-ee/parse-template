package main

import (
	"fmt"
	"github.com/lieturd/parse-template"
	"io/ioutil"
	"os"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Printf("  %s <source file> [--name=value]\n", os.Args[0])
	fmt.Println("")
	fmt.Printf("More information at https://github.com/Lieturd/parse-template\n")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	source, td := parse_template.GetTemplateData(os.Args, os.Environ())
	contents, err := ioutil.ReadFile(source)
	if err != nil {
		panic(err)
	}

	parse_template.CompileTemplate(string(contents[:]), td, os.Stdout)
}
