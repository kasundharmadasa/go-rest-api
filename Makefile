include .envrc

## build/api: build the cmd/api application
.PHONY: build/api 
build/api:
	@echo 'Building cmd/api...'
	go build -o=./bin/api ./cmd/api	

## run/api: run the cmd/api application
.PHONY: run/api 
run/api:
	go run ./cmd/api 