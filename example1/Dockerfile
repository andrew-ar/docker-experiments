# golang 1.21 was current at the time of writing, but please check for a newer version
# if using this in future to ensure you have the latest vulnerability fixes and features:
FROM golang:1.21 AS builder

# Build a static binary in the golang
COPY . /go

RUN CGO_ENABLED=0 go build -ldflags "-s" -a -installsuffix cgo -o /main main.go

# Make a new image containing the static binary and nothing else
# Set 0555 permissions (r-xr-x-r-x) on the binary
FROM scratch
COPY --from=builder --chmod=0555 /main /

# Run as a non-root user
USER 1001

# Port 8080 is to be exposed
EXPOSE 8080

# The new image will execute the main command
CMD ["/main"]