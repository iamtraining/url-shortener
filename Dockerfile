FROM golang:1.15.2-alpine3.12
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
ENTRYPOINT go run main.go