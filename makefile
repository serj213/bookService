.PHONY: run, dc, migrate

DB_URL = "postgres://postgres:root@localhost:5434/books?sslmode=disable"
PATH_MIGRATION = file://migrations

migrate:
	migrate -source ${PATH_MIGRATION} -database ${DB_URL} up

dc:
	docker-compose up --remove-orphans --build


run:
	-configPath=config/local.yaml go run cmd/main.go