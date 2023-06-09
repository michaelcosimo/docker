# Use the official Go image as the base
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o main .

# Set the command to run the application
CMD ["./main"]
