all: fk-app.proto.json fk-app.pb.go src/fk-app.pb.c src/fk-app.pb.h

node_modules/.bin/pbjs:
	npm install

fk-app.proto.json: node_modules/.bin/pbjs fk-app.proto
	pbjs fk-app.proto -t json -o fk-app.proto.json

src/fk-app.pb.c src/fk-app.pb.h: fk-app.proto
	protoc --nanopb_out=./src fk-app.proto

fk-app.pb.go: fk-app.proto
	protoc --go_out=./ fk-app.proto

clean:
