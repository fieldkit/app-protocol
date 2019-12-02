GOARCH ?= amd64
GOOS ?= linux
GO ?= env GOOS=$(GOOS) GOARCH=$(GOARCH) go
UNAME := $(shell uname)
BUILD ?= $(abspath build)
BUILDARCH ?= $(BUILD)/$(GOOS)-$(GOARCH)

all: bindings

all: bindings
	GOOS=linux GOARCH=amd64 $(MAKE) binaries-all
	GOOS=linux GOARCH=arm $(MAKE) binaries-all
	GOOS=darwin GOARCH=amd64 $(MAKE) binaries-all

install: all

bindings: fk-app.proto.json fk-app.pb.go src/fk-app.pb.c src/fk-app.pb.h org/conservify/FkApp.java

node_modules/.bin/pbjs:
	npm install

fk-app.proto.json: node_modules/.bin/pbjs fk-app.proto
	pbjs fk-app.proto -t json -o fk-app.proto.json

src/fk-app.pb.c src/fk-app.pb.h: fk-app.proto
	protoc --nanopb_out=./src fk-app.proto

fk-app.pb.go: fk-app.proto
	protoc --go_out=. fk-app.proto

org/conservify/FkApp.java: fk-app.proto
	protoc --java_out=lite:. fk-app.proto

$(BUILD):
	mkdir -p $(BUILD)

binaries-all: $(BUILDARCH)/fkdevice-cli

$(BUILDARCH)/fkdevice-cli: fkdevice-cli/*.go fkdevice/*.go
	$(GO) get ./...
	$(GO) build -o $(BUILDARCH)/fkdevice-cli fkdevice-cli/*.go

clean:
	rm -rf $(BUILD) fk-app.proto.json src/*.pb.c src/*.pb.h *.pb.go

veryclean:
