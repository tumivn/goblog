postgres:
	@echo "Starting postgres..."
	docker run --rm --name postgres -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres
	@echo "Postgres started."

createdb:
	@echo "Creating database..."
	docker exec -it postgres psql -U postgres -c "CREATE DATABASE goblog;"
	@echo "Database created."

dropdb:
	@echo "Dropping database..."
	docker exec -it postgres psql -U postgres -c "DROP DATABASE goblog;"
	@echo "Database dropped."

migrateup:
	@echo "Migrating up..."
	go run github.com/legangs/cms/cmd/dbmigrate-up
	@echo "Migrate up done."

migratedown:
	@echo "Migrating down..."
	go run github.com/legangs/cms/cmd/dbmigrate-down
	@echo "Migrate down done."