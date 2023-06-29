# Use the official Golang image as the base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download
RUN go mod tidy

# Copy the rest of the application source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Set the command to run the executable when the container starts
CMD ["./main"]
