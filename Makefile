.PHONY: web server worker build

web:
	RUNNING_MODE=web SERVER_IP=localhost:8080 PORT=80 go run main.go

server:
	RUNNING_MODE=server go run main.go

worker:
	RUNNING_MODE=worker AGGREGATOR_IP=localhost:8080 go run main.go

build:
	go build main.go