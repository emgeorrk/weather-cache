FROM golang:1.24.0-alpine3.21 as builder

RUN apk update

COPY . /weather-cache
WORKDIR /weather-cache

RUN go mod download && go build -o weather-cache

CMD ./weather-cache
