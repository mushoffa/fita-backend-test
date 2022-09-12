clean:
	rm -rf bin/* ;\
	rm -rf coverage/* ;\

generate:
	go run github.com/99designs/gqlgen generate

build-binary:
	go build -o bin/server .

build-docker:
	docker build -t mushoffa/fita-backend-test:latest .

test:
	go test -coverpkg=./... -coverprofile=coverage/coverage.out ./... ;\
	go tool cover -func=coverage/coverage.out ;\
	go tool cover -html=coverage/coverage.out ;\

run:
	go run server.go