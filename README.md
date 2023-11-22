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
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE,
  firstname VARCHAR(255) NOT NULL,
  lastname VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
```

# Build the cms app 

```bash
go build -a -installsuffix cgo -ldflags "-extldflags -static" -tags musl github.com/legangs/cms/cmd/cms
```


# Build docker image
```bash
docker build -t legangs/legangs:cms-v01 .
``` 

# Run docker container, linking with postgres container 

```bash
docker run --rm -p 8080:8080 --link postgres:postgres --name legangs-cms -e DB_HOST=postgres -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=docker -e DB_NAME=goblog legangs/legangs:cms-v01
```
