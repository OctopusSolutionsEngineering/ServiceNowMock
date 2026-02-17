# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o servicenow-mock .

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary from builder
COPY --from=builder /app/servicenow-mock .

# Copy HTML files
COPY --from=builder /app/html ./html

# Expose port 8086
EXPOSE 8086

# Run the application
CMD ["./servicenow-mock"]

