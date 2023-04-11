# Use the latest Golang image as the builder stage
FROM golang:latest as builder

# Set the working directory to /src
WORKDIR /src

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Cache dependencies before building and copying source
RUN go mod download

# Copy the source code and build the binary
COPY . .

# Build binary file
RUN GO111MODULE=on go build -ldflags="-w -s" -trimpath -o password_generator ./cmd


# Use the non-root distroless image as the runtime stage
FROM gcr.io/distroless/base:latest

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the builder stage to the runtime stage
COPY --from=builder /src/password_generator /app

# Expose port 8000 to the outside world
EXPOSE 8000

# Set the user to nonroot:nonroot
USER nonroot:nonroot

# Set the entrypoint to the binary
ENTRYPOINT ["./password_generator"]