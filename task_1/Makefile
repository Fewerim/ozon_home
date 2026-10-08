.PHONY: run
run: 
	@go run ./cmd/todo/main.go

.PHONY: check
check:
	@gofmt -l .  && \
	go build ./... && \
	go vet ./...