SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables
.DEFAULT_GOAL := help

.PHONY: $(shell egrep -o ^[a-zA-Z_-]+: $(MAKEFILE_LIST) | sed 's/://')

help:
	@echo 'Usage: make [target]'
	@echo 'Targets:'
	@egrep -o '^[a-zA-Z_-]+:' $(MAKEFILE_LIST) | perl -pe 's/^(.+):/    \1/'

generateKeys:
	go run cmd/generateKeys/main.go

SUB = "dummy_subject"

generateIdToken:
	go run cmd/generateIdToken/main.go ${SUB}
