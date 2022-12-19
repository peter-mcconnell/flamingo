FROM golang:latest

WORKDIR /src

# copy the Go script into the container
COPY main.go .

# build the Go script
RUN go build -o flamegraph main.go

RUN mkdir -p /build && cp flamegraph /build/flamegraph
