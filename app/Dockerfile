# Use the latest Alpine image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Install required packages
RUN apk add go bash bash-completion

# Copy the source Go file
COPY main.go .

# Build the Go application
RUN go build -o server main.go

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./server"]