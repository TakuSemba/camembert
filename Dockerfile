FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . $GOPATH/src/github.com/TakuSemba/camembert

WORKDIR $GOPATH/src/github.com/TakuSemba/camembert

# Install our dependencies
RUN go get github.com/tools/godep
RUN godep restore

# Install api binary globally within container
RUN go install github.com/TakuSemba/camembert

# Set binary as entrypoint
ENTRYPOINT /go/bin/camembert

# Expose default port (3000)
EXPOSE 80
