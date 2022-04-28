make build
RUNNING_MODE=web SERVER_IP=localhost:8080 PORT=80 ./main &
RUNNING_MODE=server ./main
