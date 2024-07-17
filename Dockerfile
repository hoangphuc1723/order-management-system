# Use the official Golang image as the base image
FROM golang:1.18-alpine3.16

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download all

# Copy the source code into the container
COPY . .

# List the files to ensure they are copied correctly
RUN ls -al /app

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
