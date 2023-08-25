# Use  official Go image as a parent image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to ensure dependencies are downloaded
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 3000 for the Fiber application to listen on
EXPOSE 3000

# Run the application
CMD ["./main"]