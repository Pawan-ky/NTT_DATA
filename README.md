
# Exoplanet Microservice

## Run

To run the application go inside the directory and run the command

```
  go run main.go
```

By default it will on port 8000, you can change the port in .env file.

## ENV Variables
* SERVER_PORT - used to run the application at specified port

Below paramnetrs are configurable and have their default value as given.
* Distance of planet from earth
* Planet Mass
* Planet Radius


## Building And Running Application in Docker
To build your Docker image, navigate to the directory containing your Dockerfile and run command
```
docker build -t go-app .
```

To run the docker conatiner

```
docker run -p 8000:8080 go-app
```