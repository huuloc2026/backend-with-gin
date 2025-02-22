# RESTful API with Golang and Gin

This is a simple RESTful API built with Golang and the Gin framework.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/huuloc2026/backend-with-gin
   ```

2. Install dependencies:

   ```sh
   go mod download
   ```

3. Run the application:
   ```sh
   go run main.go
   ```

## Endpoints

### Get all items

- **URL:** `/items`
- **Method:** `GET`
- **Success Response:**
  - **Code:** 200
  - **Content:** `[{ "id": 1, "name": "Item 1" }, { "id": 2, "name": "Item 2" }]`

### Get single item

- **URL:** `/items/:id`
- **Method:** `GET`
- **URL Params:** `id=[integer]`
- **Success Response:**
  - **Code:** 200
  - **Content:** `{ "id": 1, "name": "Item 1" }`

### Create new item

- **URL:** `/items`
- **Method:** `POST`
- **Data Params:** `{ "name": "New Item" }`
- **Success Response:**
  - **Code:** 201
  - **Content:** `{ "id": 3, "name": "New Item" }`

### Update item

- **URL:** `/items/:id`
- **Method:** `PUT`
- **URL Params:** `id=[integer]`
- **Data Params:** `{ "name": "Updated Item" }`
- **Success Response:**
  - **Code:** 200
  - **Content:** `{ "id": 1, "name": "Updated Item" }`

### Delete item

- **URL:** `/items/:id`
- **Method:** `DELETE`
- **URL Params:** `id=[integer]`
- **Success Response:**
  - **Code:** 204

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin Gonic](https://github.com/gin-gonic/gin) - The web framework used
- [Golang](https://golang.org/) - The programming language used
