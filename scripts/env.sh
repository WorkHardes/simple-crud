#!/bin/bash

# server
export SERVER_HOST=0.0.0.0
export SERVER_PORT=4000
export SERVER_READ_TIMEOUT=10
export SERVER_WRITE_TIMEOUT=10
export SERVER_MAX_HEADER_BYTES=1024

# docker
export DOCKERFILE_PATH=.
export DOCKERFILE_FILE_NAME=Dockerfile
export DOCKER_IMAGE_NAME=simple_crud
