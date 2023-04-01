# Inventory Management

A brand or a corporation doesn't like to have their warehouse data on the internet, so they prefer a private cluster or a single-node server for their inventory management. Therefore, I used a SQLite database for this basic GET and POST API.

## API Endpoints

## GET /BrandA

Retrieves all products in the \"BrandA\"
table.

### Request

`GET http://localhost:8000/BrandA`

### Response

```HTTP/1.1 200 OK <br/>
Content-Type: application/json

[
    {
        "name": "rice",
        "quantity": 10,
        "unit": "kg"
    }
]
```

## POST /BrandA

Creates a new product in the \"BrandA\" table.

### Request

```POST http://localhost:8000/BrandA
Content-Type:
application/json

{ <br/>*Tabspace*"name": "rice", <br/>*Tabspace*"quantity": 10, <br/>*Tabspace*"unit": "kg" }
```

### Response

HTTP/1.1 200 OK Content-Type: application/json

\[ { \"id\": 1, \"name\": \"rice\", \"quantity\": 10, \"unit\": \"kg\"
}, \... \]

## Data Model

The \"BrandA\" table has the following columns:

- `id` (integer, primary key): The unique ID of the product.
- `name` (text): The name of the product.
- `quantity` (integer): The quantity of the
  product.
- `unit` (text): The unit of measurement for the product (e.g., \"kg\", \"lbs\", etc.).

## Future Improvements

- Add support for dynamic
  creation of tables for new brands.
- Implement support for other CRUD
  operations (i.e., UPDATE and DELETE).
- Implement authentication and
  authorization for accessing the API. Add support for filtering, sorting,
  and pagination of results.
- Improve error handling and provide more
  informative error messages.
- Improve this code, we can add another POST API to create brands dynamically.
