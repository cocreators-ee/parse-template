package parse_template

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

type TemplateData struct {
	Env  map[string]string
	Args map[string]string
	Any  map[string]string
}

func GetTemplateData(args []string, env []string) (string, *TemplateData) {
	td := &TemplateData{}
	td.Env = make(map[string]string)
	td.Args = make(map[string]string)
	td.Any = make(map[string]string)

	source := args[1]

	for _, e := range env {
		pair := strings.SplitN(e, "=", 2)
		key, value := pair[0], pair[1]

		td.Env[key] = value
		td.Any[key] = value
	}

	if len(args) > 2 {
		for _, arg := range args[2:] {
			if arg[:2] != "--" {
				panic(fmt.Sprintf("Argument %s does not look like --name=value", arg))
			}

			pair := strings.SplitN(arg, "=", 2)
			key, value := pair[0][2:], pair[1]

			td.Args[key] = value
			td.Any[key] = value
		}
	}

	return source, td
}

func CompileTemplate(templateContent string, data *TemplateData, output io.Writer) {
	// TODO: Should we add .Option("missingkey=error") ? How to deal with `if .Any.local`?
	tmpl, err := template.New("template").Parse(templateContent)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(output, data)

	if err != nil {
		panic(err)
	}
}
