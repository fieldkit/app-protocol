all: fk-app.proto.json fk-app.pb.go nanopb/fk-app.pb.c nanopb/fk-app.pb.h

node_modules/.bin/pbjs:
	npm install

fk-app.proto.json: node_modules/.bin/pbjs fk-app.proto
	pbjs fk-app.proto -t json -o fk-app.proto.json

nanopb/fk-app.pb.c nanopb/fk-app.pb.h: fk-app.proto
	protoc --nanopb_out=./nanopb fk-app.proto

fk-app.pb.go: fk-app.proto
	protoc --go_out=./ fk-app.proto
