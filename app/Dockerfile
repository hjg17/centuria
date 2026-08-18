# Use an Alpine image
# In this case, it's outdated
FROM alpine:3.14


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