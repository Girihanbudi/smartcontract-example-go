FROM golang:1.19-alpine AS builder

ENV GOPATH /app
ENV GOBIN /go/bin
ENV env ${APP_ENV}

RUN mkdir -p /app/src/smartm2m-technical-test
WORKDIR /app/src/smartm2m-technical-test
ADD . /app/src/smartm2m-technical-test

RUN apk update && apk --no-cache add git openssh-client gcc g++ mercurial alpine-sdk

RUN rm -rf ./dst 
RUN CGO_ENABLED=1 go build -o ./dst/bin/app ./cmd/app/*.go

FROM alpine:3.16

LABEL Maintainer="ghanbudi@gmail.com"
LABEL Service="Smartm2m Technical Test"
LABEL Web-App="Backend"

WORKDIR /app/src/smartm2m-technical-test

RUN apk update && apk add --no-cache bash tzdata ca-certificates curl && \
    cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone \
    apk del curl tzdata && \
    rm -rf /tmp/* && \
    rm -rf /var/cache/apk/*

# Copy necessary used file
COPY --from=builder /app/src/smartm2m-technical-test/dst/bin/app /app/src/smartm2m-technical-test/main
COPY --from=builder /app/src/smartm2m-technical-test/env/config.docker.yaml /app/src/smartm2m-technical-test/env/config.yaml

EXPOSE 8081
CMD /app/src/smartm2m-technical-test/main
