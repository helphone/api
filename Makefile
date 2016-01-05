all: build

build:
	GOOS=linux GO15VENDOREXPERIMENT=1 go build -o api main.go
	docker build -t helphone/api .
	@rm ./api

build-for-test:
	@docker build -t helphone/api_test -f Dockerfile.test .

mount:
	@echo "Setup the environnement..."
	@echo "Mount the database"
	@docker run -d --name db helphone/database > /dev/null 2>&1
	@sleep 8
	@echo "Mount the importer"
	@docker run -d --name importer --env-file ./.env --link db:db helphone/importer > /dev/null 2>&1
	@sleep 5
	@echo "Launch tests"
	@-docker run --rm --name api --env-file ./.env --link db:db -p 3000:3000 helphone/api
	@docker stop importer db > /dev/null 2>&1
	@docker rm importer db > /dev/null 2>&1

up: build mount

mount-test:
	@echo "Setup the environnement..."
	@echo "Mount the database"
	@docker run -d --name db_test helphone/database
	@sleep 8
	@echo "Mount the importer"
	@docker run -d --name importer_test --env-file ./.env --link db_test:db helphone/importer
	@sleep 5
	@echo "Launch tests"
	@-docker run --rm --name api_test --env-file ./.env --link db_test:db helphone/api_test
	@docker stop importer_test db_test > /dev/null 2>&1
	@docker rm importer_test db_test > /dev/null 2>&1

test: build-for-test mount-test cleanup

cleanup:
	@echo "Cleanup in progress..."
	@-docker rm -f db db_test importer importer_test > /dev/null 2>&1 | true
