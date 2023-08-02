build:
	@go build -o ./bin/gommitter

run: build
	@./bin/gommitter

dev:
	@go run ./