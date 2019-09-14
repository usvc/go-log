run:
	$(MAKE) run_fluentd
run_fluentd:
	go run ./cmd/example-fluentd
run_fluentd_heartbeat:
	go run ./cmd/example-fluentd-heartbeat
run_json:
	go run ./cmd/example-json
run_text:
	go run ./cmd/example-text

fluentd:
	docker-compose -f ./tests/docker-compose.yml up

test:
	go test -coverprofile c.out ./...
