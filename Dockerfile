FROM golang:1.13
WORKDIR /app

RUN apt-get update
RUN apt-get install -y nodejs 
RUN go get github.com/smartystreets/goconvey gotest.tools/gotestsum
RUN curl -L https://npmjs.org/install.sh | sh

COPY src/go.* ./
RUN go mod download

