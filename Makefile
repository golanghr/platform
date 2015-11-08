NAME = platform

GO ?= go
PROTOC ?= protoc

PROTOCFLAGS = --go_out=plugins=grpc:.
PROTOSPATH = protos
PROTOS = protos/*.proto

GOPATH := $(CURDIR):$(GOPATH)

all: build

build:
	@echo "Building $(NAME) protocol buffers..."
	$(PROTOC) -I $(PROTOSPATH) $(PROTOS) $(PROTOCFLAGS)

test:
	$(GO) test
