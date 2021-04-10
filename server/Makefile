build:
	docker-compose build api-example

run:
	docker-compose up api-example

stop:
	docker-compose stop

test:
	go test -v  ./server/... -coverprofile=coverage.out
	go tool cover -html=server/coverage.out

