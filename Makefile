
all: test binaries

binaries: linux64

linux64:
	GOOS=linux GOARCH=amd64 go build -o bin/splendid cmd/main.go

deps:
	go get -u -v ./....

fmt:
	go fmt ./...

vet:
	go vet ./...

test: fmt vet
	go test -v ./...

run: linux64
	./bin/splendid

.PHONY: all deps fmt vet test run
.PHONY: binaries linux64
