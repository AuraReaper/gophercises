.PHONY: build-app app test-app-race kafka kafka-down kafka-v2 kafka-v2-down

build-app:
	@mkdir -p ./bin
	@go build -o ./bin/app ./cmd
	@chmod +x ./bin/app

app: build-app
	@./bin/app

test-app-race:
	@go clean -testcache
	@go test -race -v ./...

kafka:
	docker compose -f kafka.yaml up -d

kafka-down:
	docker compose -f kafka.yaml down

kafka-v2:
	docker compose -f kafka_v2.yaml up -d

kafka-v2-down:
	docker compose -f kafka_v2.yaml down
