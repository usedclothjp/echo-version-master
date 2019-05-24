FROM golang:1.9-alpine AS build

WORKDIR /
ADD . /go/src/gihyo/ch06/echo-version
RUN cd /go/src/gihyo/ch06/echo-version && go build -o bin/echo-version main.go

FROM alpine:3.7
ENV APP_VERSION=0.2.0
COPY --from=build /go/src/gihyo/ch06/echo-version/bin/echo-version /usr/local/bin/
COPY live.txt /
CMD ["echo-version"]
