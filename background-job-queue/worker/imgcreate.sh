#!/bin/sh

GOARCH=amd64 GOOS=linux go build -o worker

ops image create worker -c config.json -t gcp -g prod-1033 -z us-west2-a --arch=amd64
