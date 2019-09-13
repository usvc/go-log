run:
	go run ./cmd/example

fluent:
	docker-compose -f ./tests/docker-compose.yml up


build:
	go build -o ./bin/example ./cmd/example

test:
	go test ./...
