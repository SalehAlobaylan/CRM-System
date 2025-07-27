# Use Go image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy all files
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o main .

# Expose port 3000
EXPOSE 3000

# Run the application
CMD ["./main"]
