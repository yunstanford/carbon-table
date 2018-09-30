NAME:=carbon-table
MAIN:=cmd/carbon-table/main.go
MAINTAINER:="Yun Xu <yunxu1992@gmail.com>"
DESCRIPTION:="Carbon Table backend with Gin and Trie Tree"
MODULE:=github.com/yunstanford/carbon-table

GO ?= go
# export GOPATH := $(CURDIR)/_vendor
TEMPDIR:=$(shell mktemp -d)
VERSION:=$(shell sh -c 'grep "const Version" $(MAIN)  | cut -d\" -f2')

all: $(NAME)

$(NAME):
	$(GO) build $(MODULE)

test:
	$(GO) test $(MODULE)/api
	$(GO) test $(MODULE)/trie
	$(GO) test $(MODULE)/table

build:
	$(GO) build -o build/carbon-table cmd/carbon-table/main.go

version:
	echo "Version: $(VERSION)"
