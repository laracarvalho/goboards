# GoBoards

A job board API build with Go where you can create and manage your job listings.

## Features

- API ready for adding and managing job listings
  - CRUD operations
  - Pagination
- Go-Gin as router for API
- GORM and SQLite to manage persistance for the API
- Swagger for API documentation and live testing
- Makefile for managing commands and configuration
- Tests for all handlers
- Workflow CI for auditing commits

## Installation

To use this project, follow these steps:

- Clone the repository: `git clone https://github.com/laracarvalho/goboards.git`
- Install the dependencies: `go mod download`
- Build the application: `go build`
- Run the application: `./main`

After the API is running, you can use the Swagger UI to interact with the endpoints for searching, creating, editing, and deleting job opportunities. The API can be accessed at `http://localhost:$PORT/swagger/index.html`.

Default $PORT if not provided=8080.

## Dependencies

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [SQLite](https://www.sqlite.org/index.html) as the database
- [Swagger](https://swagger.io/) for API documentation and testing

## Planned features

- Dockerize and host the app
- Add authentication
- Create Candidate profiles

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

## Credits

This project was inspired by [arthur404dev](https://github.com/arthur404dev).