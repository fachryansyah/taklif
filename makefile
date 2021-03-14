run:
	go run main.go
run-cli:
	go run cli/main.go
infra-dev:
	./scripts/db.sh
build:
	go build -o taklif