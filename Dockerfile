FROM golang
ADD . /go/src/github.com/smackem/goplot
WORKDIR /go/src/github.com/smackem/goplot
RUN go install
ENTRYPOINT /go/bin/goplot
EXPOSE 9090
