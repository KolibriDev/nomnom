all: provision test build crosscompile

clean:
	rm -rf output/

provision:
	go get ./...

build:
	go build

crosscompile: build
	goxc

test:
	go test ./...