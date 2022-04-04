# Go Fiber Boilerplate with Gorm ORM
This boilerplate app is using Go version 1.17+

## Installation
- Install Golang
- do `git clone <repo>`
- cd into directory
- run `go get`

## Run Development Mode
- copy .env.example to .env and update the content
- dont forget to create the database
- do `go run main.go`
- open in your browser `localhost:3000/events`


## Run Production
```bash
docker build -t goboilerplate .
docker run -d -p 3000:3000 goboilerplate
```