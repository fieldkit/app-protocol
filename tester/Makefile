default: tester

tester: tester.go
	go build -o $@ $^

clean:
	rm -rf tester

setup:
	go get ./...
