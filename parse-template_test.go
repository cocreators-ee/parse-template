package parse_template

import (
	"bytes"
	"io"
	"log"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestGetTemplateData(t *testing.T) {
	args := []string{"./parse-template", "source.tpl"}
	env := []string{"USER=foo", "HOME=/home/foo"}

	source, td := GetTemplateData(args, env)

	Equal(t, "source.tpl", source)

	Empty(t, td.Args)

	Equal(t, "foo", td.Env["USER"])
	Equal(t, "/home/foo", td.Env["HOME"])

	Equal(t, "foo", td.Any["USER"])
	Equal(t, "/home/foo", td.Any["HOME"])
}

func TestGetTemplateDataArgs(t *testing.T) {
	args := []string{"./parse-template", "source.tpl", "--USER=bar", "--name=foobar"}
	env := []string{"USER=foo", "HOME=/home/foo"}

	_, td := GetTemplateData(args, env)

	Equal(t, "bar", td.Args["USER"])
	Equal(t, "foobar", td.Args["name"])

	Equal(t, "bar", td.Any["USER"])
	Equal(t, "foobar", td.Any["name"])
}

const TEMPLATE = `
{{- if .Any.local -}}
# Local configuration
{{- end }}

# Common configuration
ENV_USER = {{ .Env.USER }}
ARG_USER = {{ .Args.USER }}
ANY_USER = {{ .Any.USER }}
`

func TestCompileTemplate(t *testing.T) {
	args := []string{"./parse-template", "source.tpl", "--USER=bar", "--name=foobar"}
	env := []string{"USER=foo", "HOME=/home/foo"}

	_, td := GetTemplateData(args, env)

	res := new(bytes.Buffer)
	CompileTemplate(TEMPLATE, td, res)
	contentBytes, err := io.ReadAll(res)
	if err != nil {
		panic(err)
	}
	content := string(contentBytes[:])

	log.Printf("Generated template: %s", content)

	NotContains(t, content, "# Local configuration")
	Contains(t, content, "ENV_USER = foo")
	Contains(t, content, "ARG_USER = bar")
	Contains(t, content, "ANY_USER = bar")
}

func TestTemplateCondition(t *testing.T) {
	args := []string{"./parse-template", "source.tpl", "--local=1", "--USER=bar", "--name=foobar"}
	env := []string{"USER=foo", "HOME=/home/foo"}

	_, td := GetTemplateData(args, env)

	res := new(bytes.Buffer)
	CompileTemplate(TEMPLATE, td, res)
	contentBytes, err := io.ReadAll(res)
	if err != nil {
		panic(err)
	}
	content := string(contentBytes[:])

	log.Printf("Generated template: %s", content)

	Contains(t, content, "# Local configuration")
}
