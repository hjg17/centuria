# The Dockerfile for the Go application
# Build the Docker image for the Go application
FROM golang:1.26-alpine3.21 AS builder


# Set the working directory
WORKDIR /app

# Install packages
RUN apk add --no-cache bash bash-completion

# Copy the source Go file
COPY main.go .

# Build the Go application
RUN go build -o server main.go


FROM alpine:3.21

# Set the working directory
WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /app/server .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./server"]