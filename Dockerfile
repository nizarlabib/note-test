# Stage 1: build binary
FROM golang:1.23.5-alpine AS builder

WORKDIR /

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Stage 2: run binary
FROM alpine:latest

WORKDIR /

COPY --from=builder /main .
COPY .env .

EXPOSE 8080
CMD ["./main"]
