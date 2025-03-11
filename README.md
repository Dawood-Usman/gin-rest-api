# GO GIN REST API

This is a simple Book API built using Go with the Gin framework and GORM for database interactions.

## Features
- Retrieve all books
- Retrieve a single book by ID
- Create a new book
- Delete a book by ID
- Update a book by ID

## Tech Stack
- **Go** (Gin framework)
- **GORM** (ORM for database interaction)
- **PostgreSQL** (Database)
- **Docker** (Containerization)
- **Docker Compose** (Managing multi-container setup)

## Installation & Setup
### Prerequisites
- Docker & Docker Compose installed
- `.env` file configured with database credentials

### Clone the repository
```sh
$ git clone https://github.com/dawood-usman/gin-rest-api.git
$ cd gin-rest-api
```

### Setup Environment Variables By Copying `env.example` file:
   ```sh
$ cp env.example .env
```

### Run the application using Docker Compose
```sh
$ docker-compose up --build
```

This command will build and start the API along with a PostgreSQL database.

### Check running containers
```sh
$ docker ps
```

### Stopping the application
```sh
$ docker-compose down
```

### Go Version
This project is built using **Go 1.23**.

## API Endpoints

### 1. Get All Books
**Request:**
```http
GET /books
```
**Response:**
```json
[
  {
    "ID": 1,
    "CreatedAt": "2025-03-10T18:04:47Z",
    "UpdatedAt": "2025-03-10T18:04:47Z",
    "DeletedAt": null,
    "Title": "Go Programming",
    "Author": "John Doe"
  }
]
```

### 2. Get a Book by ID
**Request:**
```http
GET /books/{id}
```
**Response:**
```json
{
  "ID": 1,
  "Title": "Go Programming",
  "Author": "John Doe"
}
```

### 3. Create a Book
**Request:**
```http
POST /books
Content-Type: application/json
```
**Body:**
```json
{
  "title": "New Book",
  "author": "Jane Doe"
}
```
**Response:**
```json
{
  "ID": 2,
  "Title": "New Book",
  "Author": "Jane Doe"
}
```

### 4. Update a Book
**Request:**
```http
PUT /books/{id}
Content-Type: application/json
```
**Body:**
```json
{
  "title": "Updated Title",
  "author": "Updated Author"
}
```
**Response:**
```json
{
  "ID": 1,
  "Title": "Updated Title",
  "Author": "Updated Author"
}
```

### 5. Delete a Book
**Request:**
```http
DELETE /books/{id}
```
**Response:**
```json
{
  "message": "Book deleted successfully"
}
```

## Contributing
Feel free to fork the repository and submit a pull request.

