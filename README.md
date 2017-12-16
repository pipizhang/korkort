# korkort
Scrape Korkortonline content with Golang. This project is a exercise of scraping website with Go.

> Note: Please do not use it for persisting data and any business purpose.

[![Build Status](https://travis-ci.org/pipizhang/korkort.svg?branch=master)](https://travis-ci.org/pipizhang/korkort) [![Go Report Card](https://goreportcard.com/badge/github.com/pipizhang/korkort)](https://goreportcard.com/report/github.com/pipizhang/korkort)

## Install and run
install
```bash
go get github.com/pipizhang/korkort
```

dependencies
```bash
glide install
glide rebuild
```

run
```bash
go run main.go setup
go run main.go scrape
```

compile
```bash
bash build.sh
```

## Docker
How to build a local development environment using docker?
```bash
make docker-build
make docker-run
```
