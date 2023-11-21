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

# Build docker image
```bash
docker build -t tumivn/blog:api-v01 .
``` 

# Run docker container, linking with postgres container 

```bash
docker run --rm -p 8080:8080 --link postgres:postgres --name blog-api --env DB_HOST=postgres,DB_PORT=5432,DB_USER=postgres,DB_PASSWORD=docker,DB_NAME=goblog tumivn/blog:api-v01
```
