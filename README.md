[![Travis-CI build status](https://travis-ci.org/lieturd/parse-template.svg?branch=master)](https://travis-ci.org/lieturd/parse-template)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)


# parse-template

Render simple templates using [Golang's text/template](https://golang.org/pkg/text/template/), with environment variables and arguments as parameters to the template.

Simply put you run the tool giving it the name of a file to use as a template, and potential values to inject into the template, and you get the output to stdout.

```bash
./parse-template nginx.conf.tpl --domain=example.com > nginx.conf 
```


## Full example

The variables are injected to the template scope as follows:

 - All environment variables go under `Env`
 - All arguments go under `Args` so that `--value=foo` gets the key `value`
 - Both end up in `Any`, so arguments can override environment variables

So you can use e.g. `{{ .Env.USER }}` to print the value of `$USER` environment variable, or `{{ .Args.domain }}` for `--domain=example.com`, or `.Any.USER` or `.Any.domain` for both, allowing both environment variables and arguments to be taken into consideration.

```bash
cat <<EOF > config.yaml.tpl
{{- if .Any.LOCAL -}}
# Local configuration
foo: bar
{{- end }}

# Common configuration
user: {{ .Env.USER }}
domain: {{ .Args.domain }}
value: {{ .Any.value }}
EOF

./parse-template config.yaml.tpl --domain=example.com --value=foo > config.yaml
./parse-template config.yaml.tpl --LOCAL=1 --domain=example.local --value=bar > config-local.yaml
```

The final file contents will be:

**config.yaml**
```yaml


# Common configuration
user: yourusername
domain: example.com
value: foo
```

**config-local.yaml**
```yaml
# Local configuration
foo: bar

# Common configuration
user: yourusername
domain: example.local
value: bar
```


# License

Short answer: This software is licensed with the BSD 3-clause -license.

Long answer: The license for this software is in [LICENSE.md](./LICENSE.md), the libraries used may have varying other licenses that you need to be separately aware of.


# Financial support

This project has been made possible thanks to [Cocreators](https://cocreators.ee) and [Lietu](https://lietu.net). You can help us continue our open source work by supporting us on [Buy me a coffee](https://www.buymeacoffee.com/cocreators).

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/cocreators)
