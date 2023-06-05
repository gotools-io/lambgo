# Lambgo 

![alt text](./logo.png)

Lambgo is a free library for quickly spinning up a REST APIs with just focusing on the implementation.
I've created this for speeding up my projects creation and avoid repeting myself over and over, but decided to publish so maybe others can use it.
The concept aims to imitate the Serverless functions concept, where you can implement only a single handler function like a POST for creating your objects and scale up and down, but here you are not depending of any cloud provider, you can easily deploy this as you would like to, is just a simple golang application.
The Handler object is public so is up to you if you want to build all the crud API.
-It uses Echo from labstack to spin up the server https://github.com/labstack/echo


## How to use ?

Really simple just create a main file
Import the library
Implement the interface you need

```go
package main

import (
	lambgo "lambgo/core"
	"log"
	"os"
)

type creator struct{}

func (c creator) Create(r any) (any, error) {
	//do something with your r object
	//maybe store in the Database
	return "Something is created", nil
}

// example run
func main() {
	handler := lambgo.Handler{
		Creator: creator{},
	}
	api, err := lambgo.NewAPI(handler, "test")
	if err != nil {
		log.Fatalf("unable to configure lambgo - err: %v", err.Error())
	}
	err = api.Start(os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("unable to start lambgo - err: %v", err.Error())
	}
}
```

## How to run ?

```sh
$ go build
```

Once the build process is done
Set the environment variables 

```sh
$ SERVER_HOST=localhost SERVER_PORT=8080 ./lambgo
```

If everything is ok you should see

```sh
$ 2020/07/25 13:48:09 starting lambgo service 
$ 2020/07/25 13:48:09 sucesfully started lambgo service 

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.2
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080

```


And you can easily test your newly created API


```sh
curl --location 'http://localhost:8080/api/test' \
--header 'Content-Type: application/json' \
--data '{
    "test":"test"
}'
```

And get your results as per the example implementation

```sh
"Something is created"
```


## Execute tests ?

```sh
$ go test ./...
```


## Dockerizing your lambgo project ?

You can follow this example https://klotzandrew.com/blog/smallest-golang-docker-image

Or directly copy and paste this into your project, and you will get a really smaller docker image
How smaller? See it for yourself


```sh
REPOSITORY           TAG       IMAGE ID       CREATED          SIZE
lambgo               latest    6e74105574b0   3 seconds ago    9.67MB
```


```dockerfile
FROM golang:latest as builder

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  small-user

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /main .

USER small-user:small-user

CMD ["./main"]
```



## Feel this might be helpfull, want to help optimizing or fixing bugs ?
Happy to receive feedback, or if you think you can help improving submit a PR
