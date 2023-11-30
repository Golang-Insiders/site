help:
	@echo "Available commands:"
	@echo "  make up          : Start the database and web server"
	@echo "  make build       : Build Docker images"
	@echo "  make down        : Stop and remove Docker containers"
	@echo "  make db/migrate  : Run database migrations"
	@echo "  make db/psql     : Connect to PostgreSQL with psql"
	@echo "  make db/logs     : View Docker compose logs"
	@echo "  make help        : Show this help message"

up:
	@echo "starting db..."
	@docker compose up -d
	@sleep 1
	@echo "making migrations..."
	@goose -dir="./migrations" postgres "host=localhost port=5432 user=go-insider password=site-password dbname=db sslmode=disable" up
	@echo "starting web server..."
	@air

build:
	docker compose build

down:
	docker compose down --remove-orphans

db/migrate:
	goose -dir="./migrations" postgres "host=localhost port=5432 user=go-insider password=password dbname=db sslmode=disable" up

db/psql:
	psql "--host=localhost" "--port=5432" "--dbname=db" "--user=go-insider" "--set=sslmode=disable"

db/logs:
	docker compose logs
