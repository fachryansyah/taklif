run:
	go run main.go
run-cli:
	go run cli/main.go
infra-dev:
	echo "Creating local DB for Development"
	sudo -S cp ${PWD}/databases/data-taklif.db /etc
	echo "Local DB created!"
build:
	go build -o taklif