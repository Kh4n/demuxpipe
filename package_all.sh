#!/bin/bash
set -e

rm -rf ./bin
rm -rf ./package

env GOOS=linux GOARCH=386 go build -o bin/demuxpipe-linux-x86
env GOOS=linux GOARCH=amd64 go build -o bin/demuxpipe-linux-x64

# darwin 386 is not supported
# env GOOS=darwin GOARCH=386 go build -o bin/demuxpipe-macos-x86
env GOOS=darwin GOARCH=amd64 go build -o bin/demuxpipe-macos-x64

env GOOS=windows GOARCH=386 go build -o bin/demuxpipe-win-x86.exe
env GOOS=windows GOARCH=amd64 go build -o bin/demuxpipe-win-x64.exe

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
