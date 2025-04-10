# Filename: Makefile

# Variables
APP_NAME = mekah-tell-yuh
MAIN_FILE = cmd/web/main.go
DB_FILE = internal/db/db.go
GO_CMD = go
MIGRATE_CMD = migrate
MIGRATIONS_DIR = ./migrations

# Default target
.PHONY: all
all: run

# Run tests
.PHONY: run/tests
run/tests: vet
	$(GO_CMD) test -v ./...

# Format code
.PHONY: fmt
fmt: 
	$(GO_CMD) fmt ./...

# Vet code
.PHONY: vet
vet: fmt
	$(GO_CMD) vet ./...

# Run the application
.PHONY: run
run: vet
	$(GO_CMD) run $(MAIN_FILE) -addr=${ADDRESS} -dsn=${FEEDBACK_DB_DSN}

# Database commands
.PHONY: db/psql
db/psql:
	psql ${FEEDBACK_DB_DSN}

# Create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	$(MIGRATE_CMD) create -seq -ext=.sql -dir=$(MIGRATIONS_DIR) ${name}

# Apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running up migrations...'
	$(MIGRATE_CMD) -path=$(MIGRATIONS_DIR) -database=${FEEDBACK_DB_DSN} up

# Undo the last migration
.PHONY: db/migrations/down-1
db/migrations/down-1:
	@echo 'Running down migrations...'
	$(MIGRATE_CMD) -path=$(MIGRATIONS_DIR) -database=${FEEDBACK_DB_DSN} down 1

# Fix a SQL migration
.PHONY: db/migrations/fix
db/migrations/fix:
	@echo 'Checking migration status...'
	@$(MIGRATE_CMD) -path=$(MIGRATIONS_DIR) -database=${FEEDBACK_DB_DSN} version > /tmp/migrate_version 2>&1
	@cat /tmp/migrate_version
	@if grep -q "dirty" /tmp/migrate_version; then \
		version=$$(grep -o '[0-9]\+' /tmp/migrate_version | head -1); \
		echo "Found dirty migration at version $$version"; \
		echo "Forcing version $$version..."; \
		$(MIGRATE_CMD) -path=$(MIGRATIONS_DIR) -database=${FEEDBACK_DB_DSN} force $$version; \
		echo "Running down migration..."; \
		$(MIGRATE_CMD) -path=$(MIGRATIONS_DIR) -database=${FEEDBACK_DB_DSN} down 1; \
		echo "Running up migration..."; \
		$(MIGRATE_CMD) -path=$(MIGRATIONS_DIR) -database=${FEEDBACK_DB_DSN} up; \
	else \
		echo "No dirty migration found"; \
	fi
	@rm -f /tmp/migrate_version