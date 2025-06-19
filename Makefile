MIGRATE_CMD = go run ./cmd/migrate/main.go

migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Please provide NAME, e.g. make migrate-create NAME=create_users_table"; \
		exit 1; \
	fi
	$(MIGRATE_CMD) create $(NAME)

migrate-up:
	$(MIGRATE_CMD) up

migrate-down:
	$(MIGRATE_CMD) down

migrate-rollback:
	$(MIGRATE_CMD) steps -1

migrate-version:
	$(MIGRATE_CMD) version

migrate-force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Please provide VERSION, e.g. make force VERSION=2"; \
		exit 1; \
	fi
	$(MIGRATE_CMD) force $(VERSION)
