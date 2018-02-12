FROM alpine

WORKDIR /go/src/github.com/shimahin/gosslcommerz
COPY gosslcommerz /go/src/github.com/shimahin/gosslcommerz/

# RUN apk add --update --no-cache ca-certificates

ENV GOPATH /go

ENTRYPOINT /go/src/github.com/shimahin/gosslcommerz/gosslcommerz

# Service listens on port 80.
EXPOSE 8888