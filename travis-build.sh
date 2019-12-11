#!/usr/bin/env bash

OSES="darwin linux windows dragonfly freebsd netbsd openbsd solaris"

declare -A ARCHITECTURES
ARCHITECTURES["darwin"]="amd64 386"
ARCHITECTURES["linux"]="amd64 386 arm arm64 mips mips64 s390x"
ARCHITECTURES["windows"]="amd64 386"
ARCHITECTURES["dragonfly"]="amd64"
ARCHITECTURES["freebsd"]="amd64 386 arm"
ARCHITECTURES["netbsd"]="amd64 386"
ARCHITECTURES["openbsd"]="amd64 386 arm"
ARCHITECTURES["plan9"]="amd64 386 arm"
ARCHITECTURES["solaris"]="amd64"

# ----- END CONFIGURATION ----- #

function capitalize() {
  local text="$1"
  echo -n "${text:0:1}" | tr a-z A-Z; echo -n ${text:1:999}
}

function big_label() {
  local len
  local zero
  local fill
  local text=$(capitalize "$1")

  len=$(echo -n "$text" | wc -c)
  zero=$(printf %${len}s)
  fill=$(echo -n "$zero" | tr " " "-")

  echo
  echo
  echo
  echo  "/----$fill----\\"
  echo  "|    $zero    |"
  echo  "|    $text    |"
  echo  "|    $zero    |"
  echo "\\----$fill----/"
  echo
}

function label() {
  local len
  local zero
  local fill
  local text="$1"

  len=$(echo -n "$text" | wc -c)
  zero=$(printf %${len}s)
  fill=$(echo -n "$zero" | tr " " "-")

  echo
  echo  "/-$fill-\\"
  echo  "| $text |"
  echo "\\-$fill-/"
  echo
}

# Poor man's set -x (less spammy)
function cmd() {
  local cmd="$@"
  echo "$cmd"
  eval "$cmd"
}

# ----- END FUNCTIONS ----- #

set -e

GO111MODULE=on
CGO_ENABLED=0
GOFLAGS=-mod=vendor
ARTIFACTS="$PWD/artifacts"
mkdir -p "$ARTIFACTS"

cmd go vet . ./cmd
cmd staticcheck . ./cmd
cmd errcheck . ./cmd
cmd golangci-lint run . ./cmd

for os in $OSES; do
  big_label "$os parse-template"
  first=1

  for arch in ${ARCHITECTURES[${os}]}; do
    export GOOS="$os"
    export GOARCH="$arch"

    cd cmd

    target="parse-template-$os-$arch"
    if [[ "$os" == "windows" ]]; then
      target="$target.exe"
    fi

    label "Building $target for $os $arch"

    cmd go build -o "$target" parse-template.go
    cmd mv "$target" "$ARTIFACTS"

    cd -
  done
done
