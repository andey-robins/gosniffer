package main

import (
	"fmt"
	"os"

	"github.com/andey-robins/gosniffer/server"
	"github.com/andey-robins/gosniffer/website"
	"github.com/andey-robins/gosniffer/worker"
)

func main() {
	runMode := os.Getenv("RUNNING_MODE")

	switch runMode {
	case "web":
		startHttp()
	case "server":
		startServer()
	case "worker":
		startWorker()
	default:
		fmt.Println("Invalid environment variable 'RUNNING_MODE'")
	}
}

func startHttp() {
	website.Main()
}

func startServer() {
	server.Main()
}

func startWorker() {
	aggregatorServerAddr := os.Getenv("AGGREGATOR_IP")
	worker.Main(aggregatorServerAddr)
}
