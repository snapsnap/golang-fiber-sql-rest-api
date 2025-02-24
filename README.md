# REST API (Golang Fiber)

Create template RESTful API with Golang Fiber and SQL Database.

## 🚀 Tech Stack

- Go Fiber
- MySQL DB
- Raw SQL
- JWT (JSON Web Token)

## 📝 Features

- Clean Architecture
- Daily Logging
- JWT Authentication
- Input Validation

## 📦 Installation

```sh
git clone https://github.com/snapsnap/golang-fiber-sql-rest-api.git
cd golang-fiber-sql-rest-api
```

## ✔ Running App

Running app server :

```go
go run main.go
```

If you want to run migration & seeder, just go :

```go
go run main.go migrate
go run main.go seed
```

## ✔ Packages

If you want to install package manually, just go :

```go
// Go DotEnv
go get github.com/joho/godotenv

// MySQL Driver
// More info : https://go.dev/wiki/SQLDrivers
go get github.com/go-sql-driver/mysql

// Go Fiber
go get github.com/gofiber/fiber/v2

// JWT
go get -u github.com/gofiber/contrib/jwt
go get -u github.com/golang-jwt/jwt/v5

// Validator
go get github.com/go-playground/validator/v10

// Bcrypt
go get golang.org/x/crypto/bcrypt
```
