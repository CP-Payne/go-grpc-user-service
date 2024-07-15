# go-grpc-user-service

This project is a User Microservice that provides functionalities to retrieve user information, search users, and validate input using the decorator pattern. 

## Features
- Get a user by ID
- Get multiple users by their IDs
- Search for users based on various criteria

## Setup

1. Clone the repository
```bash
git clone https://github.com/CP-Payne/go-grpc-user-service
cd go-grpc-user-service
```

2. Install dependencies
```bash
go mod tidy
```

3. Create `.env` file and add the listening `PORT`
```bash
echo "PORT=<PORT>" > .env
```

Replace `<PORT>` with your desired port number.

4. Run the service
```bash
go run ./cmd/
```


## Usage
To use the microservice, you can interact with the gRPC endpoints defined in the handlers. Below are examples of how to call each endpoint using `grpcurl`. Once the program runs, random in-memory data will be generated to test the endpoints.

Note: You need to have `grpcurl` installed:
```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

### Example: GetUserByID

**Request:**
```bash
grpcurl -plaintext -d '{"id": 1}' localhost:3000 UserService.GetUserByID 
```
**Response:**
```json
{
  "user": {
    "id": 1,
    "firstName": "Diana",
    "lastName": "Davis",
    "city": "Columbus",
    "phone": "+1-555-7435",
    "height": 1.9589628
  }
}
```

### Example: GetUsersByIDs

#### All existing IDs:

**Request:**
```bash
grpcurl -plaintext -d '{"ids": ["1", "2"]}' localhost:3000 UserService.GetUsersByIDs 
```

**Response:**
```json
{
  "users": [
    {
      "id": 1,
      "firstName": "Diana",
      "lastName": "Davis",
      "city": "Columbus",
      "phone": "+1-555-7435",
      "height": 1.9589628
    },
    {
      "id": 2,
      "firstName": "Diana",
      "lastName": "Wilson",
      "city": "Los Angeles",
      "phone": "+1-555-4711",
      "height": 2.5208812,
      "married": true
    }
  ]
}
```

#### Some non-existing IDs provided

**Request:**
```bash
grpcurl -plaintext -d '{"ids": ["1", "999"]}' localhost:3000 UserService.GetUsersByIDs 
```

**Response:**
```json
{
  "users": [
    {
      "id": 1,
      "firstName": "Diana",
      "lastName": "Davis",
      "city": "Columbus",
      "phone": "+1-555-7435",
      "height": 1.9589628
    }
  ],
  "notFoundIds": [
    999
  ]
}
```


### Example: SearchUsers
The SearchUsers endpoint allows you to filter based on:
- First Name
- Last Name
- City
- Marriage status
- Height (Height greater than)

Note: Only provide the fields for which you want to filter on, see the below example.

Filtering based on city and marriage status:
**Request:**
```bash
grpcurl -plaintext -d '{                                                                                                                                                                                                          "city": "Columbus",
   "married": true
}' localhost:3000 UserService/SearchUsers
```

**Response:**
```json
{
  "users": [
    {
      "id": 28,
      "firstName": "John",
      "lastName": "Clark",
      "city": "Columbus",
      "phone": "+1-555-6486",
      "height": 3.3579206,
      "married": true
    },
    {
      "id": 12,
      "firstName": "Hank",
      "lastName": "Johnson",
      "city": "Columbus",
      "phone": "+1-555-3495",
      "height": 1.8366728,
      "married": true
    }
  ]
}
```

Filtering based on height:
**Request:**
```bash
grpcurl -plaintext -d '{                                                                                                                                                                                                          "height_greater_than": 1
}' localhost:3000 UserService/SearchUsers
```

**Response:**
```json
{
  "users": [
    {
      "id": 28,
      "firstName": "John",
      "lastName": "Clark",
      "city": "Columbus",
      "phone": "+1-555-6486",
      "height": 3.3579206,
      "married": true
    }
  ]
}
```
