# gosniffer
A wifi sniffing infrastructure written in golang.

## Table of Contents
- [gosniffer](#gosniffer)
  - [Table of Contents](#table-of-contents)
  - [Configuration](#configuration)
    - [Worker Mode](#worker-mode)
    - [Server Mode](#server-mode)
    - [Client Mode](#client-mode)
  - [Running Instructions](#running-instructions)
    - [Binary Creation](#binary-creation)

## Configuration

The application can be run in three different modes, detailed below. Mode selection is done by setting the environment variable `RUNNING_MODE` to one of `{worker, server, web}`

### Worker Mode

In this mode, the application listens on a wifi interface and performs all of the sniffing. It either logs data to STDOUT as it comes in or sends it to an aggregator server if an environment variable is set. Set `AGGREGATOR_IP` to the address of the aggregator server.

### Server Mode

In this mode, the application serves as the aggregator server for the worker nodes, exposes an API for querying, persists data in a database, and pushes information over websockets to clients for display.

### Client Mode

In client mode, the application opens a connection with the websocket interface of a server and serves an HTTP site with a nice display for information. Set `SERVER_IP` to the address of the server to connect to.

## Running Instructions

The program is configured with a Makefile to automatically fill in relevant environment variables. Start the applications with the following commands:
- `make web`
- `make server`
- `make worker`

### Binary Creation

To build the binary version of the program, either run `make build` or `go build main.go`. Make build is simply an alias to the go build command.