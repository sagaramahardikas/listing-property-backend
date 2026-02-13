# listing-property-backend

## Development Guide

### Prerequisite
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Go 1.25.7 or later](https://golang.org/doc/install)
- [Atlas](https://atlasgo.io/docs#installation)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Makefile](https://www.geeksforgeeks.org/installation-guide/how-to-install-make-on-ubuntu/)

## Setup
- Clone this repo, please put this repo outside of GOPATH
  ```sh
  git clone git@github.com:sagaramahardikas/listing-property-backend.git
  ```

- Copy env.sample and if necessary, modify the env value(s)
  ```sh
  cp env.sample .env
  cp dev/env.sample  dev/.env
  ```

- Download Go dependencies
  ```sh
  make dev-dep
  make dep
  ```

- Spin up dependencies with docker-compose
  ```
  make docker-dep
  ```

- Migrate the database
  ```sh
  make db-migrate
  ```

- Start the gateway server
  ```sh
  go run cmd/cli/main.go
  ```

## Local development using docker-compose

We provide `docker-compose.yaml` file to hold the non service dependencies of backend such as MySQL. To spin up the dependencies, you can do the following steps.

- Copy `dev/env.sample` to `dev/.env` and if necessary, modify the env value(s)
- Spin up the dependencies.

  ```sh
  make docker-dep
  ```

## Database Migration

- Create new migration
  ```sh
  bundle exec rake db:create_migration NAME=<migration_name>
  ```

- Up migration
  ```sh
  atlas migrate apply -u ...
  ```

- Down migration

  ```sh
  atlas migrate down -u ... --to-version ...
  ```

## Unit Test and Linter

- Run unit test

  ```sh
  make test
  ```

- Run linter

  ```sh
  make lint
  make lint-fix # to automatically fix the error
  ```
