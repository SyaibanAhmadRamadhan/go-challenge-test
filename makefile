integration-test:
	go test -v ./test/integration

unit-test:
	go test -v ./test/unit

migrate:
	go run main.go migrate

web:
	go run main.go web

build:
	go build main.go