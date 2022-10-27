FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokene/faucet
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/faucet /go/src/gitlab.com/tokene/faucet


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/faucet /usr/local/bin/faucet
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["faucet"]
