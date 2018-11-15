FROM golang:1.11

RUN mkdir -p /go/src/bigproject
WORKDIR /go/src/bigproject

ADD . /go/src/bigproject

RUN go get -v