# goblog
Learning Go by building a simple blog APIs with GoLang, Echo

# Added main dependencies

- [x] Echo  
  - `go get -u github.com/labstack/echo/v4`
  - [x] Middleware 
    - `go get -u -d github.com/labstack/echo/v4/middleware`
- [x] Postgres 
  - `go get -u -d github.com/lib/pq`
  - Ensure adding import to `main.go` file, this is an import to the concrete implementation of sql/database
    - `import _ "github.com/lib/pq"`
- [x] Dotenv 
  - `go get github.com/joho/godotenv`


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

# Create the go application 

First thing first is the `main.go` file

```go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/tumivn/goblog/cmd/handlers"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	e.Logger.Fatal(e.Start(":8080"))
}
```

For sure, this code cannot run due to the fact that we have not created the `handlers` package yet. Let's do it now.

```go
package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
```

Now, we can run the application by executing this command in your terminal

```bash
go run main.go
```
# Init the database connection

Create a new file `db.go` in the `cmd/storage` folder

```go
package storage

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err = sql.Open("postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPass, dbName, dbPort))

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
  return db
}
```


