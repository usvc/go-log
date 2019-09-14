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
	go test -coverprofile c.out ./... -v

.ssh:
	@mkdir -p ./.ssh
	@ssh-keygen -t rsa -b 4096 -f ./.ssh/id_rsa -N ""
	@cat ./.ssh/id_rsa | base64 -w 0 > ./.ssh/id_rsa.b64
