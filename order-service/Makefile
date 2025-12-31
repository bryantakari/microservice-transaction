
ifneq (,$(wildcard .env))
	include .env
	export
endif

MIGRATIONS_PATH = migrations
DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

.PHONY: migrate-up migrate-down migrate-force migrate-status migrate-create

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

migrate-force:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $(VERSION)

migrate-status:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(NAME)