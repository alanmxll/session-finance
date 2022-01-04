build-image:
	docker build -t alanmxll/finance .

run-app:
	docker-compose -f .docker/app.yml up -d

prepare-tests:
	docker-compose -f .docker/postgres.yml up -d

lint:
	golint ./...
	go fmt -n ./...

unit_test:
	go test ./...