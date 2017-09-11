
all: binaries

binaries: linux64

linux64:
	GOOS=linux GOARCH=amd64 go build -o bin/splendid cmd/main.go
