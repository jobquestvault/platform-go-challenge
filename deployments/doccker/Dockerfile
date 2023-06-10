# Start from Alpine image with Go 1.20.3
FROM golang:1.20.5-alpine3.17 AS builder

# Set the working directory in container
WORKDIR /app

# go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -o ak .

# Use a minimal image as the final base
FROM alpine:3.15

# Copy the built binary from the builder stage
COPY --from=builder /app/ak /app/ak

# Port
EXPOSE 8080

# Working directory and the command to run your application
WORKDIR /app/ak

# Entrypoint
ENTRYPOINT ["./ak"]
