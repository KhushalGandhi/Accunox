# UserService - Golang gRPC API

## Overview

The **UserService** project is a Golang-based gRPC service designed for managing user details with functionalities for fetching user details by ID, retrieving a list of users by their IDs, and searching for users based on specific criteria. This service simulates a database using an in-memory list of user details.


## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go 1.18+](https://golang.org/doc/install) (Use `go version` to check)
- [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/) (Use `protoc --version` to check)
- [protoc-gen-go](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go)
- [protoc-gen-go-grpc](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc)

### Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/KhushalGandhi/Accunox.git
    cd Accunox
    ```

2. **Install Go dependencies**:

    ```bash
    go mod tidy
    ```

3. **Install Protocol Buffers plugins**:

    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

4. **Generate Go code from .proto files**:

    ```bash
    protoc --go_out=. --go-grpc_out=. proto/user.proto
    ```

### Running the Application

1. **Build and run the application**:

    ```bash
    go run main.go
    ```

   You should see the following output if the server starts successfully:

    ```
    Server is running on port 50051
    ```

2. **Access the gRPC service**:

   The server is available on `localhost:50051`.

### gRPC Endpoints

- **Fetch User by ID**:

    ```proto
    rpc GetUserByID(GetUserByIDRequest) returns (GetUserByIDResponse);
    ```

    **Request**:
    ```json
    {
      "user_id": 1
    }
    ```

    **Response**:
    ```json
    {
      "user": {
        "id": 1,
        "fname": "Steve",
        "city": "LA",
        "phone": 1234567890,
        "height": 5.8,
        "married": true
      }
    }
    ```

- **Fetch Users by List of IDs**:

    ```proto
    rpc GetUsersByIDs(GetUsersByIDsRequest) returns (GetUsersByIDsResponse);
    ```

    **Request**:
    ```json
    {
      "user_ids": [1, 2]
    }
    ```

    **Response**:
    ```json
    {
      "users": [
        {
          "id": 1,
          "fname": "Steve",
          "city": "LA",
          "phone": 1234567890,
          "height": 5.8,
          "married": true
        },
        {
          "id": 2,
          "fname": "Jane",
          "city": "NY",
          "phone": 9876543210,
          "height": 5.5,
          "married": false
        }
      ]
    }
    ```

- **Search Users**:

    ```proto
    rpc SearchUsers(SearchUsersRequest) returns (SearchUsersResponse);
    ```

    **Request**:
    ```json
    {
      "criteria": {
        "city": "LA",
        "married": true
      }
    }
    ```

    **Response**:
    ```json
    {
      "users": [
        {
          "id": 1,
          "fname": "Steve",
          "city": "LA",
          "phone": 1234567890,
          "height": 5.8,
          "married": true
        }
      ]
    }
    ```

