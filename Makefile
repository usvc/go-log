run:
	go run ./cmd/example

build:
	go build -o ./bin/example ./cmd/example

fluent:
	docker-compose -f ./test/docker-compose.yml up
