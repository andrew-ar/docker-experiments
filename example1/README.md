# Example 1

This is a container built from scratch with a single static golang binary inside.

You can see what it builds at https://hub.docker.com/layers/aar1/hello-service/latest/images/sha256-a655bcb5f3f9aad168144cbef835413b3b03295dda98de9d2d5656998619e37e?context=explore

Notice how only the contents of the scratch container end up visible on DockerHub. The builder container layers and commands are not part of the final image.

## Usage

Example usage:

1. Set the DOCKERHUB_USER_NAME to your DockerHub user name, e.g. `export DOCKERHUB_USER_NAME=myuser`.
2. Run `make build` to build the image.
3. Run `make run` to run the image and test it, e.g. `curl http://0.0.0.0:8080/ping`
4. Run `make push` to push the image to DockerHub. This will fail if you didn't set the correct user name in step 1, or if you haven't first run `docker login`.

## Health Warning

This is a very simple demo which doesn't follow good practice, the focus is on demonstrating how layers work and what can end up inside a Docker image. Some major deviations from good practice are noted in comments.

Don't treat this as a good example of a golang application! Go and read the current golang docs, add linting and other good things.
