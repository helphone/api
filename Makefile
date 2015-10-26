all: build

build: export GOOS=linux
build:
	go build -o api main.go 

up: build
	docker-compose rm api && docker-compose build api && docker-compose up api