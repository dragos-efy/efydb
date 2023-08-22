# Start from golang base image
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN  --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags "-s -w" -o /build

FROM alpine:edge

# Copy our built binary file from the previous stage
COPY --from=builder /build /build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/build" ]
