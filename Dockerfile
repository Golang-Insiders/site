# Use the official Golang image as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod  ./
# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY cmd/ ./cmd/
COPY public/ ./public/

# Build the application
RUN go build -o /myapp cmd/site/main.go

# Expose port 3000
EXPOSE 3000

# Run the binary program
CMD ["/myapp"]
