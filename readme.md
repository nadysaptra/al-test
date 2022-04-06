# Go Fiber Boilerplate with Gorm ORM
This boilerplate app is using Go version 1.17+

## Installation
- Install Golang
- do `git clone` this repo
- cd into directory
- run `go get`

## Run Development Mode
- copy .env.example to .env and update the content
- dont forget to create the database first. Look at `.env.example` file
- do `go run main.go`. This command will auto migrate and seeder on your database.
- open in your browser `localhost:3000`
- open in your browser `localhost:3000/api-docs` for swagger-ui


## Update swagger 
- run `swag init --output docs`
- run `go run main.go` and check `localhost:3000/api-docs`

## Run Production
```bash
docker build -t goboilerplate .
docker run -d -p 3000:3000 goboilerplate
```

## List of available api
- /users
- /events
- /tickets
- /payments