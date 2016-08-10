FROM golang

ADD . /go/src/tunneltest

RUN cd /go/src/tunneltest && go get ./...
RUN go install tunneltest
WORKDIR /go/src/tunneltest
ENTRYPOINT /go/bin/tunneltest

EXPOSE 6061
