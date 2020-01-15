GOARCH ?= amd64
GOOS ?= linux
GO ?= env GOOS=$(GOOS) GOARCH=$(GOARCH) go
UNAME := $(shell uname)
BUILD ?= $(abspath build)
BUILDARCH ?= $(BUILD)/$(GOOS)-$(GOARCH)
PROTOC_VERSION = 3.11.2
PROTOC = build/bin/protoc

all: bindings

all: bindings
	GOOS=linux GOARCH=amd64 $(MAKE) binaries-all
	GOOS=linux GOARCH=arm $(MAKE) binaries-all
	GOOS=darwin GOARCH=amd64 $(MAKE) binaries-all

install: all

bindings: $(BUILD) fk-app.proto.json fk-app.pb.go src/fk-app.pb.c src/fk-app.pb.h org/conservify/FkApp.java

node_modules/.bin/pbjs:
	npm install

fk-app.proto.json: node_modules/.bin/pbjs fk-app.proto
	pbjs fk-app.proto -t json -o fk-app.proto.json

src/fk-app.pb.c src/fk-app.pb.h: fk-app.proto
	$(PROTOC) --nanopb_out=./src fk-app.proto

fk-app.pb.go: fk-app.proto
	$(PROTOC) --go_out=. fk-app.proto

org/conservify/FkApp.java: fk-app.proto
	$(PROTOC) --java_out=lite:. fk-app.proto

$(BUILD): protoc-$(PROTOC_VERSION)-linux-x86_64.zip
	mkdir -p $(BUILD)
	cd $(BUILD) && unzip ../protoc-$(PROTOC_VERSION)-linux-x86_64.zip

protoc-$(PROTOC_VERSION)-linux-x86_64.zip:
	wget "https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip"

binaries-all: $(BUILDARCH)/fkdevice-cli

$(BUILDARCH)/fkdevice-cli: fkdevice-cli/*.go fkdevice/*.go
	$(GO) get ./...
	$(GO) build -o $(BUILDARCH)/fkdevice-cli fkdevice-cli/*.go

clean:
	rm -rf $(BUILD) fk-app.proto.json src/*.pb.c src/*.pb.h *.pb.go

veryclean:
