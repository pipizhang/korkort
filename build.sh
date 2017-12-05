#!/bin/bash
OS="linux"
ARCH="amd64"
EXE="main"
GOOS=$OS GOARCH=$ARCH go build -o ${EXE} ${EXE}.go
