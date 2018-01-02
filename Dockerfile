FROM golang:1.9-alpine AS builder

ENV BIN grpcserver
ENV PKG github.com/kyeett/restful-diff-detector
ENV ARCH amd64
ARG VERSION=unknown
ENV VERSION ${VERSION}

RUN echo $VERSION
COPY . /go/src/$PKG
WORKDIR /go/src/$PKG
RUN ./build/build.sh
RUN ls /go/bin

ENTRYPOINT ["/bin/ash", "-c", "cp /go/bin/* /output"]

