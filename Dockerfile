# Stage 1: Builder
# Use golang to compile the app
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code from host to image
COPY . .

COPY ./conf/prod.conf ./conf/app.conf

# Compile
RUN CGO_ENABLED=1 GOOS=linux go build -o server .

# Stage 2: Runner
# Use a Debian-based image for better compatibility with CGO
FROM golang:1.14-trixie

# Install necessary packages for CGO and SQLite
RUN apt-get update && apt-get install -y --no-install-recommends \
    libc6 \
    libsqlite3-0 \
    sqlite3 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the rest of the application code from host to image
COPY . .
# NEW: Overwrite the app.conf with prod version
COPY ./conf/prod.conf ./conf/app.conf 

#### .........
# Ensure you have all of these COPY statements.
COPY --from=builder /app/server .
COPY --from=builder /app/static ./static/
COPY --from=builder /app/models ./models/
COPY --from=builder /app/utils ./utils/
COPY --from=builder /app/views ./views/
COPY --from=builder /app/conf/app.conf ./conf/
COPY --from=builder /app/go.* .
COPY --from=builder /app/scripts ./scripts/
COPY .env.prod .env

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
