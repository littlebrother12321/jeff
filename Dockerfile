# Stage 1: Builder
# Use golang to compile the app
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code from host to image
COPY . .

# Compile
RUN CGO_ENABLED=1 GOOS=linux go build -o server .

# Stage 2: Runner
# Use a Debian-based image for better compatibility with CGO
FROM debian:trixie-slim

# Install necessary packages for CGO and SQLite
RUN apt-get update && apt-get install -y --no-install-recommends \
    libc6 \
    libsqlite3-0 \
    sqlite3 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/static ./static/
COPY --from=builder /app/views ./views/
COPY --from=builder /app/conf/app.conf ./conf/

# Create user to run app so it's not run as root
RUN useradd appuser
# Change ownership of /app to appuser
RUN chown -R appuser:appuser /app
# Switch to appuser as the active user
USER appuser

# Command to run the application
CMD ["./server"]

# Expose the port the app runs on
EXPOSE 8080
