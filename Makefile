NAME:=carbon-table
MAINTAINER:="Yun Xu <yunxu1992@gmail.com>"
DESCRIPTION:="Carbon Table backend with Gin and Trie Tree"
MODULE:=github.com/yunstanford/carbon-table

GO ?= go
# export GOPATH := $(CURDIR)/_vendor
TEMPDIR:=$(shell mktemp -d)
VERSION:=$(shell sh -c 'grep "const Version" $(NAME).go  | cut -d\" -f2')

all: $(NAME)

$(NAME):
	$(GO) build $(MODULE)

test:
	$(GO) test $(MODULE)/trie
	$(GO) test $(MODULE)/table
