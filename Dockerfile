FROM golang:1.15-alpine

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/samtholiya/analyserService

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go build  -o /go/bin/server github.com/samtholiya/analyserService/cmd/server

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/server

# Document that the service listens on port 8080.
EXPOSE 80
