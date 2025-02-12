.PHONY: run, dc, migrate_up, migrate_down

DB_URL = "postgres://postgres:root@localhost:5434/books?sslmode=disable"
PATH_MIGRATION = file://migrations

migrate_up:
	migrate -source ${PATH_MIGRATION} -database ${DB_URL} up

migrate_down:
	@migrate -source ${PATH_MIGRATION} -database ${DB_URL} down ${VERSION}

dc:
	docker-compose up --remove-orphans --build


run:
	-configPath=config/local.yaml go run cmd/main.go