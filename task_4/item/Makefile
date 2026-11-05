DB_URL ?= postgres://postgres:postgres@localhost:5432/item?sslmode=disable

.PHONY: version
version:
	go version
	protoc --version

.PHONY: generate
generate:
	protoc -I ./ -I ./api/third_party/googleapis \
		--go_out ./pkg --go_opt paths=source_relative \
    	--go-grpc_out ./pkg --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./pkg \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--openapiv2_out ./pkg \
		--openapiv2_opt logtostderr=true \
		api/item.proto

.PHONY: test
test:
	go test ./...

.PHONY: test-api
test-api:
	go test -v -count=1 -tags=api ./test/

.PHONY: test-race
test-race:
	go test -race ./...

.PHONY: bench
bench:
	go test -bench=. -benchmem ./internal/app/

.PHONY: run
run:
	go run ./cmd/item

.PHONY: migrate-up
migrate-up:
	go tool goose -dir migrations postgres "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	go tool goose -dir migrations postgres "$(DB_URL)" down

.PHONY: migrate-status
migrate-status:
	go tool goose -dir migrations postgres "$(DB_URL)" status
