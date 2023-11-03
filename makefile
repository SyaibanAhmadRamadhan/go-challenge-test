integration-test:
	go test -v ./test/integration

migrate:
	go run main.go migrate

web:
	go run main.go web

build:
	go build main.go