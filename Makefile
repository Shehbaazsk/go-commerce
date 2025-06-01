MIGRATE_CMD = go run ./cmd/migrate/main.go

create:
	@if [ -z "$(NAME)" ]; then \
		echo "Please provide NAME, e.g. make create NAME=create_users_table"; \
		exit 1; \
	fi
	$(MIGRATE_CMD) create $(NAME)

up:
	$(MIGRATE_CMD) up

down:
	$(MIGRATE_CMD) down

rollback:
	$(MIGRATE_CMD) steps -1

version:
	$(MIGRATE_CMD) version

force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Please provide VERSION, e.g. make force VERSION=2"; \
		exit 1; \
	fi
	$(MIGRATE_CMD) force $(VERSION)
