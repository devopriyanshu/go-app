# Stage 1 - Build
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o go-app

# Stage 2 - Runtime
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/go-app .

EXPOSE 8081

CMD ["./go-app"]
