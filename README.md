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
  make start-server
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
  update schema.sql
  atlas migrate diff {migration_filename} --to file://{schema_path} --dev-url "docker://mysql/8/dev"
  ```

- Up migration
  ```sh
  atlas migrate apply -u "mysql://$(db_username):$(db_password)@$(db_host):$(db_port)/$(db_name)" --dir file://$(migration_dir)
  ```

- Down migration

  ```sh
  atlas migrate down -u "mysql://$(db_username):$(db_password)@$(db_host):$(db_port)/$(db_name)" --dir file://$(migration_dir) --to-version $(version) --dev-url "docker://mysql/8/example"
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
