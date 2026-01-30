FROM golang:1.21-alpine AS builder
RUN apk add --no-cache gcc musl-dev linux-headers

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src/

RUN CGO_ENABLED=0 GOOS=linux go build -o cutie-log src/main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/cutie-log .

RUN apk add --no-cache procps

CMD ["./cutie-log"]