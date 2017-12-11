.PHONY: build

VERSION := $(shell cat VERSION)
GITCOMMIT := $(shell git rev-parse --short HEAD)

build:
	go build -o reg

test:
	go test --cover .

dev:
	go build -o reg
	./reg ls
	./reg tags alpine
	./reg rm alpine:rm
	./reg manifest alpine:rm
	./reg ls --help
	