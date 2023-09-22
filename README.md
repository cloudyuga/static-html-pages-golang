# Static HTML Web Pages in Golang
This repository contains various static html pages that can be accessed on different paths with Golang. This is for the CloudYuga Technology Demo.

This Go application deploys two static **blue.html** and **green.html** pages at **/blue** and **/green** path respectively.
## Steps to Run the Application Locally
- Install Go, Docker and Docker Compose on the system
- Clone this Github repository
```
git clone https://github.com/cloudyuga/static-html-pages-golang.git
```
- Go inside the cloned directory and run the Go application
```
go run main.go
```
The application is deployed on port 8080 locally. Access it via the browser. You will see the following type of output.

 ![gostatic_bluepage](./gostatic_bluepage.PNG)

## Create a Docker Image of the Application
- The `Dockerfile` is present in the repository which will help in building the Docker image of the application
```
docker build -t <dockerhub-username>/go-statichtml:v1 .
```
- Run it as a container and access the application locally.
```
docker container run -d -p 8080:8080 registry.cloudyuga.guru/library/go-statichtml:v1
