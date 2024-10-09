# Movie API

A simple RESTful API to manage a collection of movies, built using Golang and Gorilla Mux.

## Features

- **GET /movies**: Retrieve all movies.
- **GET /movie/{id}**: Retrieve a specific movie by its ID.
- **POST /movies**: Add a new movie to the collection.
- **PUT /movies/{id}**: Update an existing movie by its ID.
- **DELETE /movies/{id}**: Delete a movie from the collection by its ID.

## Project Structure

The project defines the following main types and functions:

### Types

- **Movie**: A struct representing a movie with fields for ID, ISBN, Title, and Director.
- **Director**: A struct representing a director with fields for Firstname and Lastname.

### Functions

- **getMovies**: Retrieves the list of all movies.
- **getMovie**: Retrieves a single movie based on the provided movie ID.
- **createMovie**: Adds a new movie to the list.
- **updateMovie**: Updates an existing movie with new data.
- **deleteMovie**: Deletes a movie based on the provided movie ID.

## Installation

### 1. Clone the repository:

```bash
git clone https://github.com/yourusername/movie-api.git
cd movie-api
```

### 2. Install dependencies:
```bash
go mod init movie-api
go get -u github.com/gorilla/mux
```

### 3. Run the application:
```bash
go mod init movie-api
go get -u github.com/gorilla/mux
```

### The server will start at localhost:8000. (by default as defined in code)

## Usage
Use tools like Postman or curl to interact with the API.

## Example Requests
## Get all movies:

```bash
curl http://localhost:8000/movies
```

## Get a movie by ID:

```bash
curl http://localhost:8000/movie/1
```

## Create a new movie:

```bash
curl -X POST http://localhost:8000/movies -d '{"isbn": "123456", "title": "New Movie", "director": {"firstname": "Quentin", "lastname": "Tarantino"}}' -H "Content-Type: application/json"
```

## Update an existing movie:

```bash
curl -X PUT http://localhost:8000/movies/1 -d '{"isbn": "654321", "title": "Updated Movie", "director": {"firstname": "Martin", "lastname": "Scorsese"}}' -H "Content-Type: application/json"
```

## Delete a movie:

```bash
curl -X DELETE http://localhost:8000/movies/1
```

## Contributing
If you'd like to contribute to this project, feel free to submit a pull request.
