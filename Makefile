.PHONY: web server worker build

web:
	RUNNING_MODE=web go run main.go

server:
	RUNNING_MODE=server go run main.go

worker:
	RUNNING_MODE=worker go run main.go

build:
	go build main.go