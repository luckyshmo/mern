build:
	docker-compose build api-example

run:
	docker-compose up api-example

stop:
	docker-compose stop

test:
	go test -v  ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

# migrate: #alternative way to migrate DB
# 	migrate -path ./schema -database 'postgres://postgres:example@localhost/postgres?sslmode=disable' up #!wrong!!

# swag:
# 	swag init -g cmd/main.go

