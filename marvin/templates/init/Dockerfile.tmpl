{{ $go_version      := .ProjectInfo.GoVersion -}}
{{ $project_name    := .ProjectInfo.Name -}}

# Start from golang base image
FROM golang:{{$go_version}}-alpine as builder

# Install git. Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /go_server

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /opt/gofiber_server

# Copy the Pre-built binary file from the previous stage + .env file
COPY --from=builder /go_server/main .
COPY --from=builder /go_server/.env .      