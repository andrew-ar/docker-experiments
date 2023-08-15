# golang 1.21 was current at the time of writing, but please check for a newer version
# if using this in future to ensure you have the latest vulnerability fixes and features:
FROM golang:1.21 AS builder

# Some values available at build time:
ARG SECRET1
ARG BUILD_ARG
ENV BUILD_ENV=xyzzy

# This is bad practice! Don't write secrets into files in an image,
# mount them at runtime instead via k8s secrets.
RUN echo "${SECRET1}" > /secret1

# Build a static binary in the golang
COPY . /go

RUN CGO_ENABLED=0 go build -ldflags "-s" -a -installsuffix cgo -o /main main.go

# Make a new image containing the static binary and (almost) nothing else
# Set 0555 permissions (r-xr-x-r-x) on the binary
FROM scratch
COPY --from=builder --chmod=0555 /main /

# This is bad practice! Don't copy files containing a secret into an image
# mount them at runtime instead via k8s secrets.
COPY --from=builder /secret1 /

# This is also bad practice! Don't hardcode ENV values in a Dockerfile,
# inject them at runtime instead via k8s secrets:s
ENV SECRET2=3OulHF7XmLnPuexH/SEkmurtRabtz/vWWQuI
# (this isn't a real secret, it's a randomly generated value)

# Run as a non-root user
USER 1001

# Port 8080 is to be exposed
EXPOSE 8080

# The new image will execute the main command
CMD ["/main"]