package main

import (
	"github.com/lieturd/parse-template"
	"io/ioutil"
	"os"
)

func main() {
	source, td := parse_template.GetTemplateData(os.Args, os.Environ())
	contents, err := ioutil.ReadFile(source)
	if err != nil {
		panic(err)
	}

	parse_template.CompileTemplate(string(contents[:]), td, os.Stdout)
}
