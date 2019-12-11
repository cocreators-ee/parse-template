{{- if .Any.LOCAL -}}
# Local configuration
foo: bar
{{- end }}

# Common configuration
user: {{ .Env.USER }}
domain: {{ .Args.domain }}
value: {{ .Any.value }}
