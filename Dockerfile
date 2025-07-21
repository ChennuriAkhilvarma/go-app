# Use the official Golang image with required version
FROM golang:1.24.3 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go app
RUN go build -o app

# Use a minimal base image for the final binary
FROM debian:bullseye-slim

WORKDIR /root/
COPY --from=builder /app/app .

CMD ["./app"]
