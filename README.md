# Book Management API Documentation

## Overview
The Book Management API allows users to manage a collection of books. It provides endpoints to create, read, update, and delete books from the collection.

## Base URL
```
http://localhost:9010/v1
```

## Endpoints

### Get All Books
```
GET /book
```
**Response:**
- `200 OK` on success
- Returns a list of all books

### Get a Single Book
```
GET /book/{id}
```
**Parameters:**
- `id` (required): The ID of the book to retrieve

**Response:**
- `200 OK` on success
- Returns the details of the specified book
- `404 Not Found` if the book does not exist

### Create a New Book
```
POST /book
```
**Request Body:**
- `name`: The title of the book
- `author`: The author of the book
- `publication`: The book's publication

**Response:**
- `201 Created` on success
- Returns the created book
- `400 Bad Request` if the request body is invalid

### Update an Existing Book
```
PUT /book/{id}
```
**Parameters:**
- `id` (required): The ID of the book to update

**Request Body:**
- `name`: The title of the book
- `author`: The author of the book
- `publication`: The book's publication

**Response:**
- `200 OK` on success
- Returns the updated book
- `400 Bad Request` if the request body is invalid
- `404 Not Found` if the book does not exist

### Delete a Book
```
DELETE /book/{id}
```
**Parameters:**
- `id` (required): The ID of the book to delete

**Response:**
- `200 OK` on success
- `404 Not Found` if the book does not exist

## Error Handling
The API uses standard HTTP status codes to indicate the success or failure of an API request. The response body contains a JSON object with an `error` field describing the error.

## Example Requests
### Get All Books
```bash
curl -X GET http://localhost:9010/v1/book
```

### Get a Single Book
```bash
curl -X GET http://localhost:9010/v1/book/1
```

### Create a New Book
```bash
curl -X POST http://localhost:9010/v1/book -H "Content-Type: application/json" -d '{"title": "New Book", "author": "Author Name"}'
```

### Update an Existing Book
```bash
curl -X PUT http://localhost:9010/v1/book/1 -H "Content-Type: application/json" -d '{"title": "Updated Title"}'
```

### Delete a Book
```bash
curl -X DELETE http://localhost:9010/v1/book/1
```

## Rate Limiter Middleware

To prevent abuse and ensure fair usage, the API includes a rate limiter middleware. Each user is allowed a requests per second. If the limit is exceeded, the API will respond with a `429 Too Many Requests` status code.

**Response:**
- `429 Too Many Requests` if the rate limit is exceeded

**Headers:**
- `Retry-After`: The number of seconds to wait before making a new request
