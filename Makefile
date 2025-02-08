include .env
export

MIGRATE := migrate -path migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DB_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

.PHONY: migrate-create migrate-up migrate-down migrate-force migrate-version db-start

# Start database
db-start:
	docker-compose up -d

# Create a new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $${name}

# Run all up migrations
migrate-up:
	$(MIGRATE) up

# Rollback one migration
migrate-down:
	$(MIGRATE) down 1

# Rollback all migrations
migrate-down-all:
	$(MIGRATE) down

# Force set migration version
migrate-force:
	@read -p "Enter version to force: " version; \
	$(MIGRATE) force $${version}

# Show current migration version
migrate-version:
	$(MIGRATE) version 