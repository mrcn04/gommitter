build:
	@go build -o ./bin/gommitter

run: build
	@./bin/gommitter

dev:
	@go run ./

build-docker:
	@docker build --build-arg GH_KEY=<your_key>  --build-arg GH_USERNAME=<your_username> --build-arg REPO_NAME=<your_repo_name> -t gommitter .

run-docker:
	@docker run -e GH_USERNAME=<your_username> -e REPO_NAME=<your_repo_name> -e GH_KEY=<your_key>  gommitter
	