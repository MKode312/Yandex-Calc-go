# Distributed calculator with persistence and multi-user mode

## Description

The user sends his login and password via http to the address http://localhost:8080/api/v1/register, then sends the same data to http://localhost:8080/api/v1/login. In response, the user receives a unique JWT token, which is written into the user's cookie. Then, after authentication and authorization, the user can send arithmetic expressions to http://localhost:8080/api/v1/calculate, get a list of all sent expressions from the database by sending a request to http://localhost:8080/api/v1/expressions, and delete each sent expression from the list and the database by sending a request to http://localhost:8080/api/v1/expression/{id}.

## Strcuture of the project

```
├── cmd
│ ├── agent
│ │ └── main.go
│ └── orchestrator
│ │ └── main.go
├── database
│ └── storage.db    #sqlite3
├── internal
│ ├── grpc
│ │ ├── agent
│ │ │ └── agent.go
│ │ └── orchestrator
│ │ │ └── orchestrator.go
│ ├── http
│ │ ├── handlers
│ │ │ ├── auth
│ │ │ │ └── auth.go
│ │ │ ├── expression
│ │ │ │ └── expression.go
│ ├── storage
│ │ ├── expression_storage.go
│ │ ├── storage.go
│ │ └── user_storage.go
│ ├── utils
│ │ ├── agent
│ │ │ ├── calculation
│ │ │ │ ├── calculation.go
│ │ │ │ └── stack.go
│ │ │ ├── infix_to_postfix
│ │ │ │ ├── infix_to_postfix.go
│ │ │ │ └── stack.go
│ │ │ ├── validator
│ │ │ │ └── validator.go
│ │ ├── orchestrator
│ │ │ ├── jwts
│ │ │ │ └── jwts.go
│ │ │ ├── manager
│ │ │ │ └── manager.go
├── protos
| | ├── gen
| | | ├── agent.pb.go
| | | └── agent_grpc.pb.go
| | ├── proto
│ | | └── agent.proto
├── go.mod
├── go.sum
└── README.md
```

### REST API requests

| POST /api/v1/register | -> | RegisterUserHandler | -> | RegisterUser |

| POST /api/v1/login | -> | LoginUserHandler | -> | LoginUser |

| POST /api/v1/calculate | -> | CreateExpressionHandler | -> | InsertExpression |

| GET /api/v1/expressions | -> | GetExpressionsHandler | -> | SelectExpressionsByID |

| DELETE /api/v1/expression/{id} | -> | DeleteExpressionHandler | -> | DeleteExpression |

## Requirements

- Go 1.24.2
- Supported operations: +, -, *, /
- Operator precedence and parentheses
- Parallel execution of operations

## Installation

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/MKode312/Yandex-Calc-go.git
    ```
2. **Go to the project directory**:
    ```bash
    cd Yandex-Calc-go
    ```
3. **Install all required dependencies**:
    ```bash
    go mod tidy

---

# Running the System

## 1. Starting the Orchestrator

```bash
go run ./cmd/orchestrator/main.go
```

## 2. Starting the Agent

```bash
go run ./cmd/agent/main.go
```

## API Endpoints

### 1. Registration

```bash
POST /api/v1/register
```

Example request:

```bash
curl --location 'http://localhost:8080/api/v1/register' \
--header 'Content-Type: application/json' \
--data '{
  "login": "exampleLogin",
  "password": "examplePassword"
}'
```

Successful response (201):
"You have successfully registered"

### 2. Login

```bash
POST /api/v1/login
```

Send a request with the same body:

```bash
curl --location 'http://localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
  "login": "exampleLogin",
  "password": "examplePassword"
}'
```

Example response (200):

```json
{
  "token": <your_example_JWT_token>
}
```

### 3. Sending the expression

```bash
POST /api/v1/calculate
```

Example request:

```bash
curl -b "auth_token=<your_example_JWT_token>" \
     -H "Content-Type: application/json" \
     -d '{"expression":"2+2*2"}' \
     http://localhost:8080/api/v1/calculate
```

Successful response (201):

```json
{
    "id":1
}
```

### 4. Get all sent expressions

```bash
GET /api/v1/expressions
```

Example request:

```bash
curl -b "auth_token=<your_example_JWT_token>" http://localhost:8080/api/v1/expressions
```

Successful response (200):

```json
[
    {
        "id": 1,
        "expression": "2+2*2",
        "result": "6",
        "date": "2025/04/30 10:07:52",
        "status": "done"
    }
]
```

### 5. Delete an expression by id

```bash
DELETE /api/v1/expression/{id}
```

Example request:

 ```bash
curl -X DELETE -b "auth_token=<your_example_JWT_token>" http://localhost:8080/api/v1/expression/1
```
Successful response (202):

```json
{
    "Expression with this id was successfully deleted": 1
}
```


## Scenario Examples

### 1. Starting the System

#### 1.1 Starting the Orchestrator

First, you need to start the orchestrator. Open a terminal and run the following command:

```bash
go run ./cmd/orchestrator/main.go
```

#### 1.2 Starting the Agent

In another terminal, start the agent:

```bash
go run ./cmd/agent/main.go
```

Now your system is ready to work.

### 2. Using the API

##### Successful Scenario

To register a new user, execute the following request:

```bash
curl --location 'http://localhost:8080/api/v1/register' \
--header 'Content-Type: application/json' \
--data '{
  "login": "exampleLogin",
  "password": "examplePassword"
}'
```

**Successful Response (201)**: User successfully registered.

##### Unsuccessful Scenario

If you try to register a user with an already existing login, you will receive an error:

```bash
curl --location 'http://localhost:8080/api/v1/register' \
--header 'Content-Type: application/json' \
--data '{
  "login": "exampleLogin",
  "password": "examplePassword"
}'
```

**Unsuccessful Response (400)**:
"This login is already registered"

#### 2.2 Logging In

##### Successful Scenario

After registration, you can log in by sending the same request with your login and password:

```bash
curl --location 'http://localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
  "login": "exampleLogin",
  "password": "examplePassword"
}'
```

**Example Response (200)**:

```json
{
  "token": "<your_example_JWT_token>"
}
```

##### Unsuccessful Scenario

If you enter an incorrect login or password, you will receive an error:

```bash
curl --location 'http://localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
  "login": "wrongLogin",
  "password": "wrongPassword"
}'
```

**Unsuccessful Response (401)**:
"Invalid login or password"

#### 2.3 Sending an Expression for Calculation

##### Successful Scenario

Now you can send a mathematical expression for calculation:

```bash
curl -b "auth_token=<your_example_JWT_token>" \
     -H "Content-Type: application/json" \
     -d '{"expression":"2+2*2"}' \
     http://localhost:8080/api/v1/calculate
```

**Successful Response (201)**:

```json
{
    "id": 1
}
```

##### Unsuccessful Scenario

If you provide an invalid expression, you will receive an error:

```bash
curl -b "auth_token=<your_example_JWT_token>" \
     -H "Content-Type: application/json" \
     -d '{"expression":"invalid_expression"}' \
     http://localhost:8080/api/v1/calculate
```

**Unsuccessful Response (400)**:
"Invalid expression"

#### 2.4 Retrieving All Submitted Expressions

##### Successful Scenario

To get a list of all submitted expressions, execute the following request:

```bash
curl -b "auth_token=<your_example_JWT_token>" http://localhost:8080/api/v1/expressions
```

**Successful Response (200)**:

```json
[
    {
        "id": 1,
        "expression": "2+2*2",
        "result": "6",
        "date": "2025/04/30 10:07:52",
        "status": "done"
    }
]
```

##### Unsuccessful Scenario

If the authentication token is invalid or missing, you will receive an error:

```bash
curl -b "auth_token=invalid_token" http://localhost:8080/api/v1/expressions
```

**Unsuccessful Response (403)**:
"Unauthorized"

#### 2.5 Deleting an Expression by ID

##### Successful Scenario

If you need to delete an expression by its ID, use the following request:

```bash
curl -X DELETE -b "auth_token=<your_example_JWT_token>" http://localhost:8080/api/v1/expression/1
```

**Successful Response (202)**:

```json
{
    "Expression with this id was successfully deleted": 1
}
```

##### Unsuccessful Scenario

If you try to delete an expression with a non-existent ID, you will receive an error:

```bash
curl -X DELETE -b "auth_token=<your_example_JWT_token>" http://localhost:8080/api/v1/expression/9999
```

**Unsuccessful Response (404)**:
"Expression not found"
