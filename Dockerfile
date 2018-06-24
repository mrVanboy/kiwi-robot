FROM golang:1.10 AS build-env
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /bin/dep
ADD ./src /go/src
RUN chmod +x /bin/dep && \
    cd /go/src/kiwi-robot && \
    dep ensure -vendor-only && \
    go test ./... && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goapp

# final stage
FROM scratch
MAINTAINER Ivan Boyarkin
WORKDIR /app
COPY --from=build-env /go/src/kiwi-robot/goapp /app
ENTRYPOINT ["./goapp"]