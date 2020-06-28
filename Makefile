watch:
	docker-compose run -p 8080:8080 build goconvey

test:
	docker-compose run build go test ./...

build:
	docker-compose run build go build -o bin/main handlers/commandline/main.go

run:
	docker-compose run build go run ./handlers/commandline/main.go
