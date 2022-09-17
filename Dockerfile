FROM golang:1.19.1 AS build

WORKDIR /goapp

COPY ./server ./server
COPY ./lib ./lib

WORKDIR /goapp/server
RUN go mod download

RUN go build -o /tmp/goserver main.go

FROM debian:11.5

WORKDIR /

COPY --from=build /tmp/goserver /goserver

EXPOSE 8080
ENTRYPOINT [ "/goserver" ]
