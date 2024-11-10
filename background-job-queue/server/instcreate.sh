#!/bin/sh

ops instance create server -c config.json -t gcp -g prod-1033 -p 8080 -z us-west2-a
