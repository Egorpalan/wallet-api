build:
	docker-compose build wallet-api

run:
	docker-compose up wallet-api

test:
	go test -v ./...