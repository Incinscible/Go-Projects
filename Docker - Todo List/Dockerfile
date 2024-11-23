FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o todolist main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/todolist .
EXPOSE 8080
CMD ["./todolist"]
