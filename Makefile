help:
	@echo "Commands:";
	@echo "build-docker-image: builds docker image with name '$(DOCKER_IMAGE_NAME)'";
	@echo "run-docker-container: runs container '$(DOCKER_IMAGE_NAME)'"

build-docker-image:
	docker build -t $(DOCKER_IMAGE_NAME) $(DOCKERFILE_PATH)

run-docker-container:
	docker run -it $(DOCKER_IMAGE_NAME)
