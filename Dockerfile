# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build statically linked binary (disable CGO)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Stage 2: Minimal runtime image using distroless
FROM gcr.io/distroless/static-debian11

# Set non-root user (typically UID 65532 for distroless)
USER nonroot:nonroot

# Copy the binary from the builder
COPY --from=builder /app/server /

# Use an unprivileged port if possible (like 8080)
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/server"]
