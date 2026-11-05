-include .env
export

DB_URL = postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DBNAME)?sslmode=disable

.PHONY: migrate-up
migrate-up:
	@goose -dir migrations postgres "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	@goose -dir migrations postgres "$(DB_URL)" down

.PHONY: migrate-status
migrate-status:
	@goose -dir migrations postgres "$(DB_URL)" status
