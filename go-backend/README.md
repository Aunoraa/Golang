# Go Backend 

A Go (Golang) Backend Clean Architecture project with Gin, MongoDB, JWT Authentication Middleware, Test, and Docker.



## About me



You can connect with me on:

- [LinkedIn](https://www.linkedin.com/in/le-tuan-b70055254/)
- [GitHub](https://github.com/Aunoraa)
- lekhactuan1412@gmail.com
- 0916369594


## Major Packages used in this project


- Check packages in `go.mod`.


### How to run this project?
#### Run without Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install `go` if not installed on your machine.
- Install `MongoDB` if not installed on your machine.
- Important: Change the `DB_HOST` to `localhost` (`DB_HOST=localhost`) in `.env` configuration file. `DB_HOST=mongodb` is needed only when you run with Docker.
- Run `go run cmd/main.go`.
- Access API using `http://localhost:8080`

#### Run with Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install Docker and Docker Compose.
- Run `docker-compose up -d`.
- Access API using `http://localhost:8080`

### How to run the test?

```bash
# Run all tests
go test ./...
```

### How to generate the mock code?

In this project, to test, we need to generate mock code for the use-case, repository, and database.

```bash
# Generate mock code for the usecase and repository
mockery --dir=domain --output=domain/mocks --outpkg=mocks --all

# Generate mock code for the database
mockery --dir=mongo --output=mongo/mocks --outpkg=mocks --all
```

### SWAGGER UI
```azure
http://localhost:8080/swagger/index.html
```

Whenever you make changes in the interfaces of these use-cases, repositories, or databases, you need to run the corresponding command to regenerate the mock code for testing.







