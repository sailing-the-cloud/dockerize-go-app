# Compile source code.
FROM golang:1.14-stretch as builder
# You can change SRC_DIR
ENV SRC_DIR="/go/src/github.com/sailing-the-cloud/dockerize-go-app"
WORKDIR $SRC_DIR
COPY go.mod .
COPY go.sum .
# Copy all source and build it.
# This layer will be rebuilt whenever a file has changed in the source directory.
COPY ./ .
RUN GOOS=linux GOPROXY=https://proxy.golang.org go build -v -a -mod=vendor -o /bin/server .

# Build final image.
FROM frolvlad/alpine-glibc
RUN apk update && \
    apk add --no-cache ca-certificates \
    openssl \
    curl \
  && update-ca-certificates \
  && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /bin/server server
ENTRYPOINT ["./server"]