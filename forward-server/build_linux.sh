#!/bin/sh

export GoDevWork=/Users/tavenli/Desktop/Work/port-forward-v2

echo "Build For Linux..."
export GOOS=linux
export GOARCH=amd64
export GOPATH=${GoDevWork}:${GOPATH}
go build -o forward-server

echo "--------- Build For Linux Success!"


