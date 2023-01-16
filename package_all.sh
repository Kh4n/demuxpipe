#!/bin/bash
set -e

rm -rf ./bin
rm -rf ./package

GOOS=linux GOARCH=386; go build -o bin/demuxpipe-$GOOS-$GOARCH
GOOS=linux GOARCH=amd64; go build -o bin/demuxpipe-$GOOS-$GOARCH

# darwing 386 is not supported
# GOOS=darwin GOARCH=386; go build -o bin/demuxpipe-$GOOS-$GOARCH
GOOS=darwin GOARCH=amd64; go build -o bin/demuxpipe-$GOOS-$GOARCH

GOOS=windows GOARCH=386; go build -o bin/demuxpipe-$GOOS-$GOARCH.exe
GOOS=windows GOARCH=amd64; go build -o bin/demuxpipe-$GOOS-$GOARCH.exe

mkdir package

for file in bin/*.exe
do
    zip package/$(basename "${file%.*}").zip $file LICENSE
    rm -f $file
done
for file in bin/*
do
    tar -czvf package/$(basename "$file").tar.gz $file LICENSE
    rm -f $file
done
