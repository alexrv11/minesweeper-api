# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.12.7-alpine

# Add Maintainer Info
LABEL maintainer="Alex Ventura <alex.rv11@gmail.com>"

RUN mkdir -p /go/src/deviget/minesweeper-api

# Set the Current Working Directory inside the container
WORKDIR /go/src/deviget/minesweeper-api

# Copy go mod and sum files
COPY . /go/src/deviget/minesweeper-api


# Build the Go app
RUN go install /go/src/deviget/minesweeper-api

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/go/bin/minesweeper-api"]