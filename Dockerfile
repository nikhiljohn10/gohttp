FROM golang:1.19.1 AS build

WORKDIR /goapp

COPY ./server ./server
COPY ./lib ./lib
COPY ./go.mod .
RUN go mod download

COPY main.go .

RUN go build -o /tmp/goserver main.go

FROM debian:11.5

WORKDIR /

COPY --from=build /tmp/goserver /usr/bin/gohttp

EXPOSE 8080
ENTRYPOINT [ "gohttp" ]
