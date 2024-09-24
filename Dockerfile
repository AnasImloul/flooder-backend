# Use the official Golang image as the build stage
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first (for dependency caching)
COPY go.mod go.sum ./

# Download dependencies
RUN GOPROXY=direct go mod download

# Copy the entire project, including subdirectories and files
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o "flooder" ./cmd/flooder

# Use a minimal image for the final container
FROM alpine:latest

# Set the working directory in the minimal container
WORKDIR /app

# Copy the binary and all necessary files from the builder container
COPY --from=builder /app /app

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the application
CMD ["./flood-backend"]
