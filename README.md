# Fetch Receipt Processor

This is the exercise for the Fetch backend developer position.  https://github.com/fetch-rewards/receipt-processor-challenge

## How to start the service
This service requires a Golang environment 1.23.3 or above. There are multiple ways to start it. The service listens to port 8080.

### Start locally
Check out the project on a device with the proper Golang environment. Start the service by calling ``go run main.go`` on ./ directory. The service will start and take requests in a few seconds.  

This project uses two external modules. Use the following commands to get the modules if the local device doesn't have them. 

```
go get -u github.com/gin-gonic/gin

go get -u github.com/google/uuid
```

### Build a docker image and start
The project has a Dockerfile to build an image and run on a Docker environment. The following steps are needed to start the service with this path.
```
//Build the Docker image
docker build -t fetch-receipt-processor

//Check if the image was built successfully 
docker images

//Run the service 
docker run -p 8080:8080 fetch-take-home
```
These should be enough to start the service and take requests. 

### Built images on the repository
A built image of this project has been shared with a public repo on Docker Hub. Pulling the image and then running it should start the service too. https://hub.docker.com/r/xh50/fetch-receipt-processor/tags
