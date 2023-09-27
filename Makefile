build:
	@go build -o ./bin/gommitter

run: build
	@./bin/gommitter

dev:
	@go run ./

build-docker:
	@docker build --build-arg GH_KEY=<your_key> -t gommitter .

run-docker:
	@docker run -e GH_KEY=<your_key>  gommitter
	