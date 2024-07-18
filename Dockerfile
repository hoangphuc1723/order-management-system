# Use an official Golang runtime as a parent image
FROM golang:1.18

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download necessary Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go build`
CMD ["./main"]
