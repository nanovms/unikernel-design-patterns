#!/bin/sh

GOARCH=amd64 GOOS=linux go build -o server

ops image create server -c config.json -t gcp -g prod-1033 -z us-west2-a --arch=amd64
