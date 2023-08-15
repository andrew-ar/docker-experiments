# Example 3

This is a container built from scratch with a single static golang binary inside.

You can see what it builds at https://hub.docker.com/layers/aar1/nicer-hash-service/latest/images/sha256-be32d1e2060d81fc753f28be5272994c2679f5850ed6ec4f4bab4fc534718eb3?context=repo

Notice how only the contents of the scratch container end up visible on DockerHub. The builder container layers and commands are not part of the final image.

This is intended to be a slightly less-awful version of example 2, without such glaring security issues. It's still not perfect though. Tools and practices evolve over time, it's likely this will become outdated so keeping up with the latest news is always a good idea.

## Usage

Example usage:

1. Set the DOCKERHUB_USER_NAME to your DockerHub user name, e.g. `export DOCKERHUB_USER_NAME=myuser`.
2. Set a couple of random secret values for the service to use, e.g. `export SECRET1=$(openssl rand -base64 21); export SECRET2=$(openssl rand -base64 21);`
3. Run `make build` to build the image.
4. Run `make run` to run the image and test it, e.g. `curl http://0.0.0.0:8080/ping?parameter=12345`. **The output is based on a hash of the URL parameters and some secrets, so try different parameter values to get different outputs.**
5. Run `make push` to push the image to DockerHub. This will fail if you didn't set the correct user name in step 1, or if you haven't first run `docker login`.

## Health Warning

Although this is the least-worst example here, it's still a very simple demo which doesn't follow every good practice, the focus is on demonstrating how layers work and what can end up inside a Docker image.

Don't treat this as a good example of a golang application! Go and read the current golang docs, add linting and other good practices.
