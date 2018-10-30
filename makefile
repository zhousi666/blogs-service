all: build

export PROJECT_NAME=blogs-service
export PROJECT_DIR=$(shell pwd)
export BUILD_PATH=$(PROJECT_NAME)/cmd

export PROJECT_SOURCE_DIR=$(PROJECT_DIR)/$(PROJECT_NAME)
export PROJECT_SOURCE_DIR_DOCKER=/go/src/$(PROJECT_NAME)

export PROJECT_BUILD_PATH=$(PROJECT_DIR)/build
export PROJECT_BUILD_PATH_DOCKER=/go/build

export GOLANG_VERSION=golang:1.10.4
export TAG=0.0.1

build:
	sudo docker run --rm -v $(PROJECT_SOURCE_DIR):$(PROJECT_SOURCE_DIR_DOCKER)	\
		-v $(PROJECT_BUILD_PATH):$(PROJECT_BUILD_PATH_DOCKER) $(GOLANG_VERSION)	\
		go build -ldflags "-X main.Version=$(TAG)" -o $(PROJECT_BUILD_PATH_DOCKER)/$(PROJECT_NAME) $(BUILD_PATH)