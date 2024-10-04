## Technology

- Golang `go 1.15 or up`
- PostgresSQL
- Redis

## How to Start Develop

Before developing this project, you need to do some setup

1. Setup .env.yml add value from your local source
2. Run database migrations Follow the database migration section below

## Database Migration

I'm using migrations tools [goose](https://github.com/pressly/goose)

```$command
// Check status of the migrations
$ goose -dir migrations postgres "user=postgres dbname=yourdatabasename password=password sslmode=disable" status

// Create database migrations file
$ goose -dir migrations create create_something_table sql

// Up the migrations
$ goose -dir migrations postgres "user=postgres dbname=yourdatabasename password=password sslmode=disable" up

// Down the migrations by one
$ goose -dir migrations postgres "user=postgres dbname=yourdatabasename password=password sslmode=disable" down

// Reset all the migrations (Will erase all the migrations on database)
$ goose -dir migrations postgres "user=postgres dbname=yourdatabasename password=password sslmode=disable" reset
```

## Install All Package

```$command
$ make install
```

## Run HTTP API

```$command
$ make run
```

But when there is any new feature, please register the path into `Makefile`
