# go-listing

go-listing is an simple app built with golang to create product & get product list.

## Features
- Create product
- Get Product list
- Simple Data Cache with Redis
- Run in Docker

## Architecture
- go-listing uses a clean architecture referenced by go-clean-arch `https://github.com/bxcodec/go-clean-arch`. It has 4 domain layers such
  - Models Layer
  - Repository Layer
  - Usecase Layer
  - Delivery Layer

- Here's some reason why referencing go-clean-arch:
  - Easily testable:
      - Domains are easly and independently testable.
  - Maintainable
      - A way more organized code. Even with separated layers, still encapsulates logic business that may not affect other logic when having updates/changes.
  - Easier to understand.
      - Code should not depend on one developer. Readable codes will saves time and effort of developers in future.

## Stacks

- Golang
- Postgresql
- Redis
- Docker
- Makefile `brew install make`

## Clone

First clone this repo by run:

```sh
$ git clone git@github.com:notblessy/mini-wallet.git
```

## Run go-listing in local
### Init

Firstly, run:

```sh
$ go mod tidy
```

### Environment

- The sample environments are provided in root folder
  - If you run go-listing in local, use `env.sample` to be `.env` file.

### Database Migration

- Ensure you have already installed `Makefile` and created the database. To migrate tables, run:

```sh
$ make migration
```

### Running app

- To run HTTP server, hit:
```sh
$ make run
```

## Run go-listing in Docker
### Environment

- The sample environments are provided in root folder
  1. If you run go-listing in Docker, use `env.Docker.sample` to be `.env` file.

### Build

Then, build image with:

```sh
$ docker-compose build
```

### Database Migration

- Ensure you have already created the database. To migrate tables, ssh to container by:

```sh
$ docker exec -it GO_LISTING_CONTAINER_ID sh
```

- Replace `GO_LISTING_CONTAINER_ID` to container id of golisting. You should be in the folder /app, then run migration with:
```sh
$ make migration
```

### Running app

- To run app, hit:
```sh
$ docker-compose up
```

#
## API Documentations

- To test the API, import `postman collection` from folder `api-docs/`. All the API is available there.

- If you're not using `postman` here are the available API endpoint to test
  - `GET: localhost:8181/products?sort=name-DESC`
    - Request can be sorted by query `?sort=` and by value `name` or `price` with order type `ASC` or `DESC` separated with dash (`-`).
  - `POST: localhost:8181/products`
    - Create request needs a body payload:
```json
# Create Request Payload
{
  "name": "Oud Vibrant Leather",
  "price": 399999,
  "description": "A premium perfume by Zara",
  "quantity": 23
}
```

#

## Author

```
I Komang Frederich Blessy
```
