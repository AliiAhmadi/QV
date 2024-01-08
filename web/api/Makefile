# Ali Ahmadi (C) 2024

.DEFAULT_GOAL := help

## help: print this help message
.PHONY: help
help:
	@echo "Usage:"
	@sed -n "s/^##//p" ${MAKEFILE_LIST} | column -t -s ':' | sed -e "s/^/ /"


## run: run the web application
.PHONY: run
run:
	@go run .