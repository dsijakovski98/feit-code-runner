# Stage 1: Build the Go application
FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd parsers/go && go install
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/feit-code-runner

# Stage 2: Build the JS/TS parsers (unchanged)
FROM oven/bun:latest AS js-ts-builder

WORKDIR /app/parsers/js-ts

COPY parsers/js-ts/package.json parsers/js-ts/bun.lockb ./
RUN bun install

# Stage 3: Final image
FROM alpine:3.18

# Install runtime dependencies (crucial!)
RUN apk add --no-cache ca-certificates libc6-compat

WORKDIR /app

# Copy only the necessary artifacts
COPY --from=builder /app/bin/feit-code-runner /app/bin/feit-code-runner
COPY --from=js-ts-builder /app/parsers/js-ts /app/parsers/js-ts

RUN chmod +x /app/bin/feit-code-runner

EXPOSE 4242
CMD ["./bin/feit-code-runner"]
