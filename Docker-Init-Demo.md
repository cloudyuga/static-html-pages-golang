# For Docker init Demo with Docker Desktop
- Install Docker Desktop 4.18 or higher according to your operating system from [here](https://www.docker.com/products/docker-desktop/).

 - Clone this GitHub repository with `docker-init-demo` branch.
```
https://github.com/cloudyuga/static-html-pages-golang.git -b docker-init-demo
```
- Go inside the cloned directory and run the Go application
```
go run main.go
```
- Execute `docker-init` command and follow the instructions according to your application platform.Read more about it [here](https://collabnix.com/introducing-docker-init-generating-docker-assets-for-your-projects/)
  It will generate the following docker assets with template content according to the application platform you have chosen
1. .dockerignore
2. Dockerfile
3. compose.yaml for Docker Compose 

- Update the `Dockerfile` content with the below content to create a docker image for this Go application.
```
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
```
- Build the Docker image
```
docker build -t <dockerhub-username>/go-statichtml:v1 .
```
- Run it as a container and access the application locally.
```
docker container run -d -p 8080:8080 registry.cloudyuga.guru/library/go-statichtml:v1
```
