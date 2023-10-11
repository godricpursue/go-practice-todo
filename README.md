# First Go Project - TODO

## Overview

This is a Go project that demonstrates connecting to a PostgreSQL database using GORM, performing CRUD (Create, Read, Update, Delete) operations on a `TodoInput` table, and managing sensitive database information with environment variables.

## Prerequisites

Before you can run this project, make sure you have the following prerequisites:

- Go programming language installed: https://golang.org/
- PostgreSQL database installed and running
- PostgreSQL driver for GORM installed:
```
go get -u gorm.io/gorm
go get gorm.io/driver/postgres
```
- godotenv package installed: `go get github.com/joho/godotenv`

## Setup

1. Clone this repository to your local machine:

```shell
   git clone https://github.com/godricpursue/go-practice-todo.git
   cd go-practice-todo
```
2. Create a .env file in the project directory and set the following environment variables with your database connection information:
```shell
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=your_database_password
   DB_NAME=your_database_name
   DB_PORT=5432
```
3. Load environment variables using the following command:
```shell
   go run main.go
```
## Usage
The main code is located in the main.go file.
Functions for adding, deleting, and toggling to-do items are in the functions.go file.
You can uncomment the function calls in main.go to test the different database operations (adding, deleting, and toggling) on the TodoInput table.
