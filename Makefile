project = $(shell basename $(shell pwd))

test:
	go test -failfast -p=1 -count=1 ./...

lint:
	test -z $(shell go fmt ./...)
	go vet ./...
	staticcheck

testAll: lint test

dockerUp:
	docker-compose up

dockerDown:
	docker-compose down

run:
	./... go run ./main.go