# docker-experiments

A series of simple experiments to demonstrate the relationship between a Dockerfile and the built image.

Example usage:

1. Set the DOCKER_USER_NAME to your Docker Hub user name, e.g. `DOCKER_USER_NAME=myuser`.
2. Run `make build` to build the image.
3. Run `make run` to run the image and test it, e.g. `curl http://0.0.0.0:8080/ping`
4. Run `make push` to push the image to Docker Hub. This will fail if you didn't set the correct user name in step 1, or if you haven't first run `docker login`.
