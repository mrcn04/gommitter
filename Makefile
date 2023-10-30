build:
	@go build -o ./bin/gommitter

run: build
	@./bin/gommitter

dev:
	@go run ./

build-docker:
	@docker build --tag gommitter-gcp --platform linux/amd64 .

run-docker:
	@docker run -p 8081:8081 gommitter-gcp
	