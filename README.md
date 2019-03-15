# multi-stage
Multi-stage builds in Docker. The final webserver image is 5.33MB compared to the 317MB golang:alpine image.

## Build
```
docker build -t webserver .
```

## Run
```
docker run -p 8000:8000 webserver
```
