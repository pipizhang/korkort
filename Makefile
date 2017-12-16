SHELL := /bin/bash

.PHONY: help test docker-build docker-run docker-clean

help: ## This help message
	@echo "usage: make [target]"
	@echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m: \2/')"

test: ## Run tests
	exec go test -v ./korkort

docker-build: ## Build docker image
	@echo "build docker image ..."
	@exec docker build --no-cache -t korkort .
	@echo "done"

docker-run: ## Start a container
	@echo "start docker container ..."
	@exec docker run --rm -it --name "korkort-dev" --hostname "korkort-dev" -v $$(pwd):/go/src/github.com/pipizhang/korkort korkort

docker-clean: ## Run docker prune
	@echo "y" | docker image prune

