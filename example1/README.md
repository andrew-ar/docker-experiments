# Example 1

This is a container built from scratch with a single static golang binary inside.

You can see what it builds at https://hub.docker.com/layers/aar1/hello-service/latest/images/sha256-0bcadb722f8576763651bde6c2a8a980c916068f746e7ebe433c50d4cfdc4e1d?context=explore

Notice how only the contents of the scratch container end up visible on Docker Hub. The builder container layers and commands are not part of the final image.

## Health Warning

This is a very simple demo which doesn't follow every good practice, the focus is on demonstrating how layers work and what can end up inside a Docker image.

Don't treat this as a good example of building a Docker golang container! Go and read something like https://snyk.io/blog/containerizing-go-applications-with-docker/ for a better example.
