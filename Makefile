DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=microwave

.PHONY: up deps check-db run

up: deps check-db run

deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download

check-db:
	@echo "Checking if the PostgreSQL database is running..."
	@pg_isready -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME)
	@if [ $$? -ne 0 ]; then \
		echo "Database is not available. Please start PostgreSQL."; \
		exit 1; \
	fi
	@echo "Database is running."

run:
	@echo "Running the project..."
	@go run ./cmd/main.go
