export MIGRATIONS_DIR=migrations
export USER=user
export PASSWORD=password
export NAME=db
export HOST=localhost
export PORT=5432

# install goose
.PHONY: install
install: 
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: docker-up
docker-up: docker compose up -d

.PHONY: docker-down
docker-down: docker compose down

.PHONY: up
up:
	goose -dir "$(MIGRATIONS_DIR)" postgres "user=$(USER) password=$(PASSWORD) dbname=$(NAME) host=$(HOST) port=$(PORT) sslmode=disable" up

.PHONY: down
down:
	goose -dir "$(MIGRATIONS_DIR)" postgres "user=$(USER) password=$(PASSWORD) dbname=$(NAME) host=$(HOST) port=$(PORT) sslmode=disable" down

.PHONY: status
status:
	goose -dir "$(MIGRATIONS_DIR)" postgres "user=$(USER) password=$(PASSWORD) dbname=$(NAME) host=$(HOST) port=$(PORT) sslmode=disable" status

.PHONY: all
all:
	print "Hello world"