clean:
	rm -rf bin/*

generate:
	go run github.com/99designs/gqlgen generate

build:
	go build -o bin/server .

run:
	go run server.go