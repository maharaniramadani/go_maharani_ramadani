#builder
FROM golang:1.17.7-alpine3.14 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o api main.go

#runer
FROM alpine:3.14

COPY --from=builder /app/api /app/

WORKDIR /app

EXPOSE 8080

CMD ["./api"]