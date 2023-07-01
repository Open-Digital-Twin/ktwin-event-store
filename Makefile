build:
	go build -o ./

run:
	go run .

swag:
	swag init

wire:
	wire ./internal/app && \
	wire ./internal/app/context/...