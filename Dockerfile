# Use the official Golang image as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod  ./
# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o /site ./cmd/site/

# Expose port 3000
EXPOSE 3000

# Run the binary program
CMD ["/site"]
