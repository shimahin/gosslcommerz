FROM golang:1.8
MAINTAINER fahim <fahim.shoumik@gmail.com>
ADD . /go/src/github.com/shimahin/gosslcommerz
WORKDIR /go/src/github.com/shimahin/gosslcommerz
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo .

ENV GOPATH /go

ENTRYPOINT /go/src/github.com/shimahin/gosslcommerz/gosslcommerz

# Service listens on port 8080.
EXPOSE 8080