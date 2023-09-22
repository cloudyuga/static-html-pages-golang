# syntax=docker/dockerfile:1

FROM golang:1.20-alpine AS build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod  ./
RUN go mod download

COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o  /godocker
EXPOSE 8080
CMD ["/godocker"]