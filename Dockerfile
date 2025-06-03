FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -o main main.go

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/main /app
COPY .env.example .env
CMD ["./main"]