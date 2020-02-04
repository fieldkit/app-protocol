GOARCH ?= amd64
GOOS ?= linux
GO ?= env GOOS=$(GOOS) GOARCH=$(GOARCH) go
UNAME := $(shell uname)
BUILD ?= $(abspath build)
BUILDARCH ?= $(BUILD)/$(GOOS)-$(GOARCH)
PROTOC_VERSION = 3.11.2
PROTOC_BIN = build/bin
PROTOC = $(PROTOC_BIN)/protoc
PROTO_NAME = fk-app
JAVA_DEP = org/fieldkit/app/pb/FkApp.java

all: bindings
	GOOS=linux GOARCH=amd64 $(MAKE) binaries-all
	GOOS=linux GOARCH=arm $(MAKE) binaries-all
	GOOS=darwin GOARCH=amd64 $(MAKE) binaries-all

install: all

binaries-all: $(BUILDARCH)/fkdevice-cli

$(BUILDARCH)/fkdevice-cli: fkdevice-cli/*.go fkdevice/*.go
	$(GO) get ./...
	$(GO) build -o $(BUILDARCH)/fkdevice-cli fkdevice-cli/*.go

bindings: $(PROTO_NAME).proto.json $(PROTO_NAME).pb.go src/$(PROTO_NAME).pb.c src/$(PROTO_NAME).pb.h $(JAVA_DEP)

$(PROTO_NAME).proto.json: build $(PROTO_NAME).proto
	node_modules/.bin/pbjs $(PROTO_NAME).proto -t json -o $(PROTO_NAME).proto.json

src/$(PROTO_NAME).pb.c src/$(PROTO_NAME).pb.h: build $(PROTO_NAME).proto
	PATH=$(PATH):$(PROTOC_BIN) $(PROTOC) --plugin=protoc-gen-nanopb=build/nanopb/generator/protoc-gen-nanopb --nanopb_out=./src $(PROTO_NAME).proto

$(PROTO_NAME).pb.go: build $(PROTO_NAME).proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	$(PROTOC) --go_out=./ $(PROTO_NAME).proto

$(JAVA_DEP): build $(PROTO_NAME).proto
	$(PROTOC) --java_out=./ $(PROTO_NAME).proto

build: protoc-$(PROTOC_VERSION)-linux-x86_64.zip
	mkdir -p build
	cd build && unzip ../protoc-$(PROTOC_VERSION)-linux-x86_64.zip
	git clone https://github.com/nanopb/nanopb.git build/nanopb
	pip install protobuf
	npm install

protoc-$(PROTOC_VERSION)-linux-x86_64.zip:
	wget "https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip"

veryclean: clean

clean:
	rm -rf build node_modules *.pb.go *.pb.c *.pb.h $(PROTO_NAME).proto.json *.pb.go org
