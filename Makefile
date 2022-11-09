help:
	@echo "Commands:";
	@echo "start-server: starts server at host '$(SERVER_HOST)' and port '$(SERVER_PORT)'";
	@echo "lint: lint all code with golangci-lint";
	@echo "pre-commit: runs pre-commit hooks";
	@echo "init-pre-commit: install pre-commit into .git";
	@echo "build-docker-image: builds docker image with name '$(DOCKER_IMAGE_NAME)'";
	@echo "run-docker-container: runs container '$(DOCKER_IMAGE_NAME)'"

start-server:
	go run ./cmd/app/main.go

lint:
	golangci-lint run

init-pre-commit:
	pre-commit install

pre-commit:
	pre-commit run --all-files

build-docker-image:
	docker build -t $(DOCKER_IMAGE_NAME) $(DOCKERFILE_PATH)

run-docker-container:
	docker run -it $(DOCKER_IMAGE_NAME)
