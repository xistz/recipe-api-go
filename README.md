# Recipe API Go

![Recipe API Go workflow](https://github.com/xistz/recipe-api-go/workflows/Recipe%20API%20Go%20workflow/badge.svg)

A sample CRUD API for recipes. Built with Golang and MySQL

Recipe API can be accessed at <https://recipe-api-go.herokuapp.com/>

## Available Scripts

In the project directory, you can run:

### `docker-compose build`

Builds the development container

### `docker-compose up`

Starts development server

### `docker-compose run --rm api go test -v ./...`

Runs tests

## Required ENV variables

### `DB_USER`

Username used to login to the MySQL DB

### `DB_PASSWORD`

Password used to login to the MySQL DB

### `DB_ADDRESS`

Address(including port) of the MySQL DB

### `DB_NAME`

Name of the MySQL DB to use for this service
