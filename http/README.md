# http

basic http server in golang for listing and creating sale orders using grpc and nats streaming.

## Prerequisites

This project requires

- [Go 1.16+](https://golang.org/)
- make

## Getting Started

### Installation

Run `make` to install dependencies.

```bash
make install
```

### Usage

Setup `.env` file according to template provided in `.env.sample`, then run locally using

```bash
make serve
```

## Testing

Open your browser at [http://localhost:50042](http://localhost:50042) and start sending request.

We can do a `GET` request to `/saleorders`

The server returns the following response:

```json
// http://localhost:50042/saleorders
[
  {
    "id": 3,
    "email": "davidbuchanan@northwind.com",
    "payment_method": "cash",
    "products": [
      {
        "name": "Canned Coffee",
        "quantity": 10
      }
    ]
  },
  {
    "id": 2,
    "email": "nancydavolio@northwind.com",
    "payment_method": "cashless",
    "products": [
      {
        "name": "Coca Cola",
        "quantity": 5
      },
      {
        "name": "Mint Candy",
        "quantity": 200
      }
    ]
  },
  {
    "id": 1,
    "email": "michaelsuyama@northwind.com",
    "payment_method": "cash",
    "products": [
      {
        "name": "Chicken",
        "quantity": 1
      },
      {
        "name": "Pepsi",
        "quantity": 3
      },
      {
        "name": "Momogi",
        "quantity": 10
      }
    ]
  }
]
```

[Swagger documentation](./../api/openapi-spec/swagger.yaml) can be found inside `documentation` [folder](../api/openapi-spec)
