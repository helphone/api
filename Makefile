all: build
up: build mount
test: mount-test cleanup

create-builder:
	@echo "The builder is in construction..."
	@docker build -t helphone/api-builder -f scripts/Dockerfile.build .

build:
	@echo "API build started..."
	@docker run --name builder -v $$(pwd):/go/src/github.com/helphone/api helphone/api-builder
	@echo "Build finished, the final docker image is in construction"
	docker build -t helphone/api -f scripts/Dockerfile .
	@docker rm builder
	@rm api
	@echo "Image construction and cleanup are finished"

mount:
	@echo "Setup the environnement..."
	@echo "Mount the database"
	@docker run -d --name db_api helphone/database > /dev/null 2>&1
	@sleep 10
	@echo "Mount the importer"
	@docker run -d --name importer_api --env DB_USERNAME=postgres --env DB_PASSWORD=postgres --link db_api:db helphone/importer > /dev/null 2>&1
	@sleep 5
	@echo "Start the API"
	-docker run --rm --name api --env DB_USERNAME=postgres --env DB_PASSWORD=postgres --link db_api:db -p 3000:3000 helphone/api

mount-test:
	@echo "Setup the environnement..."
	@echo "Mount the database"
	@docker run -d --name db_api_test helphone/database
	@sleep 10
	@echo "Mount the importer"
	@docker run -d --name importer_api_test --env DB_USERNAME=postgres --env DB_PASSWORD=postgres --link db_api_test:db helphone/importer
	@sleep 5
	@echo "Launch tests"
	-docker run -it --rm --env DB_USERNAME=postgres --env DB_PASSWORD=postgres --link db_api_test:db -v $$(pwd):/go/src/github.com/helphone/api helphone/api-builder /bin/sh -c "./scripts/test.sh"

cleanup:
	@echo "Cleanup in progress..."
	@-docker rm -f db_api db_api_test importer_api importer_api_test > /dev/null 2>&1 | true
