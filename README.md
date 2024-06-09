# goblog
Learning Go by building a simple blog APIs with GoLang, Echo

# Install make on Windows 

```bash
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
choco install make
```

#Install go-migrate cli

On Windows, you can install the cli using scoop

```bash
irm get.scoop.sh | iex
scoop install go-migrate
```

On Mac, you can install the cli using brew

```bash
brew install golang-migrate
```

On Linux, you can install the cli using apt

```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

# Create Database 
Run postgres on docker 

```bash
make postgres
```

Create database 

```bash
make createdb
```
Migrate database to the latest version

```bash
make migrateup
```

If you want to migrate to a go back one version, use this command

```bash
make migratedown
```
Just in case you want to drop the database because it is dirty or something

```bash
make dropdb
```

To have more flexibility, you can use the migrate cli directly

# run the app

```bash
go run github.com/tumivn/goblog/cmd/api
``` 

# Build the cms app 

```bash
go build -a -installsuffix cgo -ldflags "-extldflags -static" -tags musl go build -a -installsuffix cgo -ldflags "-extldflags -static" -tags musl github.com/tumivn/goblog/cmd/api
```


# Build docker image
```bash
docker build -t tumivn/tumivn:goblog-v01 .
``` 

# Run docker container, linking with postgres container 

```bash
docker run --rm -p 8080:8080 --link postgres:postgres --name legangs-auth -e DB_HOST=postgres -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=docker -e DB_NAME=goblog -e PORT=8080 -e JWT_SECRET=my_secret_key tuminv/tumivn:auth-v01
```

# Database migration
Using the [golang-migrate/migrate](https://github.com/golang-migrate/migrate) library apply for database migration 

Create seperated handler and entry point to run migration.
