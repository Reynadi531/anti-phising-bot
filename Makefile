build:
	go build -o bin/main main.go

run:
	go run main.go

docker-dev:
	docker-compose -f docker/docker-compose.dev.yaml up

docker-build:
	docker build . -f docker/Dockerfile -t phising-checker:latest

dev:
	air
