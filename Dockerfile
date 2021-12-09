FROM golang:1.14.6-alpine as builder

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app main.go

FROM alpine:latest

WORKDIR /api
COPY templates /api/templates
COPY --from=builder /api/app .

EXPOSE 8080

CMD ["./app"]