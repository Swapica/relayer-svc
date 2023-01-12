FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/Swapica/relayer-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/relayer-svc /go/src/github.com/Swapica/relayer-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/relayer-svc /usr/local/bin/relayer-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["relayer-svc"]
