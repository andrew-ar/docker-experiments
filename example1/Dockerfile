FROM golang:1.21 AS builder

# Build a static binary in the golang
COPY . /go

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o /main main.go

# Make a new image containing the static binary and nothing else
FROM scratch
COPY --from=builder /main /

# The new image will execute the main command
CMD ["/main"]