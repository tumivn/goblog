# goblog
Learning Go by building a simple blog APIs with GoLang, Echo

# Create Database and Main Tables
Execute this command in your terminal to run postgres on docker 

```bash
docker run --rm --name postgres -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres
```

Create database and tables

```sql
CREATE DATABASE goblog;

CREATE TABLE users (
  id uuid DEFAULT uuid_generate_v4 (),
  username VARCHAR(255) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE,
  firstname VARCHAR(255) NOT NULL,
  lastname VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
```

# Run the cms app 

```bash
go run github.com/legangs/auth/cmd/auth
```

# Build the cms app 

```bash
go build -a -installsuffix cgo -ldflags "-extldflags -static" -tags musl github.com/legangs/auth/cmd/auth
```


# Build docker image
```bash
docker build -t legangs/legangs:auth-v01 .
``` 

# Run docker container, linking with postgres container 

```bash
docker run --rm -p 8080:8080 --link postgres:postgres --name legangs-auth -e DB_HOST=postgres -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=docker -e DB_NAME=goblog -e PORT=8080 -e JWT_SECRET=my_secret_key legangs/legangs:auth-v01
```

# Database migration
Using the [golang-migrate/migrate](https://github.com/golang-migrate/migrate) library apply for database migration 

Create seperated handler and entry point to run migration.
