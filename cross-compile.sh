#!/bin/bash
GOARCH=amd64 GOOS=linux CGO_ENABLED=1 GOEXPERIMENT=arenas CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ go build -ldflags "-linkmode external -extldflags -static" -o api main.go