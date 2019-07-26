# go-data-converter-service 

Exposes an API written in `Go` to convert data from JSON to binary using `protobuf`.

## Getting started

Compile `service.proto` file:

    protoc --go_out=plugins=grpc:. src\proto\service.proto

## Usage

### JSON -> binary

POST http://localhost:8080/api/convert

Content-Type: `application/json`

Body:

    {
      "body": "value"
    }