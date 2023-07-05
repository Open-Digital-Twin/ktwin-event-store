build:
	go build -o ./

run-local:
	export LOCAL=true && \
	go run .

run:
	go run .

swag:
	swag init

wire:
	wire ./internal/app && \
	wire ./internal/app/context/...