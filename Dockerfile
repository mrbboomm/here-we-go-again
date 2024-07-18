# Use an official Go runtime as a parent image
FROM golang:1.18-alpine

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files first
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Install dependencies and build the Go application
RUN go mod tidy
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Run the Go application
CMD ["./main"]
