# Build PlatON in a stock Go builder container
FROM golang:1.21-alpine3.19 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers g++ llvm bash git gmp-dev openssl-dev
ADD AppChain-SDK /AppChain-SDK
RUN cd /AppChain-SDK/simapp && go build
RUN cd /AppChain-SDK/tools/cmd && go build -o tool

# Pull PlatON into a second stage deploy alpine container
FROM alpine:latest
RUN apk add --no-cache curl ca-certificates libstdc++ bash tzdata gmp-dev

WORKDIR /cluster

COPY --from=builder /AppChain-SDK/simapp/simapp /usr/local/bin/
COPY --from=builder /AppChain-SDK/tools/cmd/tool /usr/local/bin/
ADD AppChain-SDK/simapp/docker/quickstart/cluster /cluster

