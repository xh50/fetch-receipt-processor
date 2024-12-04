# Build application
FROM golang:1.23 as builder

# Set the working dir
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN  go mod download

# Copy code
COPY . .
RUN go build -o app .

# create a minimal runtime image
FROM debian:bookworm-slim

# Set the working directory
WORKDIR /root/

# Copy the binary from builder stage 
COPY --from=builder ./app .

EXPOSE 8080

CMD ["./app"]