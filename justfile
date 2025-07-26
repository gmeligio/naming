set shell := ["powershell.exe", "-c"]

default:
  @just --justfile {{justfile()}} --list

# Run tests
test: lint
	go test ./... -v -timeout 120m

# Run and fix static checks
lint:
	golangci-lint run --fix
