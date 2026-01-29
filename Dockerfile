# ---- Build stage ----

# Uses the official Go 1.25 image based on Alpine Linux
# AS builder names this stage so we can refer to it later
FROM golang:1.25-alpine AS builder

# Sets /app as the working directory inside the container
WORKDIR /app

# Copies only go.mod and go.sum into the container
# Why only these files? This is a Docker caching optimization:
# -> Go dependencies only change when go.mod or go.sum changes
# -> Docker will reuse this layer if your source code changes but deps don’t
COPY go.mod go.sum ./

# Downloads all Go dependencies listed in go.mod
# Stores them in the module cache inside the container
RUN go mod download

# Copies the rest of the project files into /app
COPY . .

# CGO_ENABLED=0: Disables C bindings → produces a fully static binary
# GOOS=linux: Target OS is Linux
# GOARCH=amd64: Target CPU architecture

# go build -o ginger-root ./cmd/server -> builds the Go package located at ./cmd/server and
# Outputs a binary named ginger-root - a single Linux executable that does NOT need Go installed to run
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ginger-root ./cmd/server


# ---- Run stage ----

# Starts a new, clean image - no Go, no compiler, no build tools, just a minimal Linux system
FROM alpine:latest

# Sets /app as the working directory
WORKDIR /app

# Copies only the compiled binary from the build stage
# Ignores all source code, dependencies, and build tools
COPY --from=builder /app/ginger-root .

# Documents that the container listens on port 8080
# Does not actually open the port
EXPOSE 8080

# Sets the default command when the container starts (runs automatically when the container is launched)
# Runs the compiled Go server
CMD ["./ginger-root"]
