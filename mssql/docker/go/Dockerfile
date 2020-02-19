FROM golang:latest

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

RUN apt-get update
RUN apt-get upgrade git

#ENV GIT_CURL_VERBOSE 0

#RUN git config --global http.sslVerify false
#RUN git config --global http.proxyAuthMethod basic

# build directories
RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app

WORKDIR /go/src/app

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Build my app
#RUN go build -o /app/main .
CMD ["/app/main"]