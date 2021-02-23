#!/bin/bash

#https://medium.com/travis-on-docker/how-to-cross-compile-go-programs-using-docker-beaa102a316d
#https://dh1tw.de/2019/12/cross-compiling-golang-cgo-projects/

docker run --rm -v "$GOPATH":/go -w /go/src/github.com/exitstop/speaker exitstop/golang_bakend_msys2:latest \
    /bin/sh -c '

export PATH="$PATH:/usr/include/:/mingw64/include:/mingw64/bin"
export LIBRARY_PATH=$LIBRARY_PATH:/usr/lib/:/mingw64/lib/
set -ex
GOOS=windows GOOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 go build -v -o build/speaker.exe cmd/voice/main.go
set -e
'

#docker run --rm -it -v "$GOPATH":/go -w /go/src/github.com/exitstop/speaker golang:1.4.2-cross sh -c '
#rm -rf /usr/local/go
#wget https://golang.org/dl/go1.16.linux-amd64.tar.gz --no-check-certificate
#tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz
#export PATH=$PATH:/usr/local/go/bin

#for GOOS in windows; do
  #for GOARCH in amd64; do
    #echo "Building $GOOS-$GOARCH"
    #export GOOS=$GOOS
    #export GOARCH=$GOARCH
    #go mod download
    #go build -o bin/ironcli-$GOOS-$GOARCH cmd/voice/main.go
  #done
#done
#'

#docker run --rm -it -v "$GOPATH":/go -w /go/src/github.com/iron-io/ironcli golang:1.4.2-cross sh -c '
#for GOOS in darwin linux windows; do
  #for GOARCH in 386 amd64; do
    #echo "Building $GOOS-$GOARCH"
    #export GOOS=$GOOS
    #export GOARCH=$GOARCH
    #go build -o bin/ironcli-$GOOS-$GOARCH
  #done
#done
#'
