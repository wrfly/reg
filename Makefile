.PHONY: build test dev

NAME = reg
PKG = github.com/wrfly/$(NAME)

VERSION := $(shell cat VERSION)
COMMITID := $(shell git rev-parse --short HEAD)
BUILDAT := $(shell date +%Y-%m-%d)

CTIMEVAR = -X $(PKG)/version.CommitID=$(COMMITID) \
	-X $(PKG)/version.Version=$(VERSION) \
	-X $(PKG)/version.BuildAt=$(BUILDAT)
GO_LDFLAGS = -ldflags "-w $(CTIMEVAR)"
GO_LDFLAGS_STATIC = -ldflags "-w $(CTIMEVAR) -extldflags -static"

build:
	go build -tags "$(BUILDTAGS)" $(GO_LDFLAGS) -o $(NAME) .

test:
	go test --cover .

dev: build
	./reg help
# ./reg ls
# ./reg tags alpine
# ./reg rm alpine:rm
# ./reg manifest alpine:rm
# ./reg ls --help
	