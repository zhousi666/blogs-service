#go 默认$GOPATH根本为 ``/go/src/`` , go编译器会自动找到 ``blogs-service/cmd`` 下的 main

# Go parameters
PROJECT_NAME=blogs-service
BINARY_NAME=blogs
GOLANG_VERSION=golang:1.10.4
TAG=0.0.1

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

PROJECT_DIR=$(shell pwd)
BUILD_PATH=$(PROJECT_NAME)/cmd
PROJECT_SOURCE_DIR=$(PROJECT_DIR)/$(PROJECT_NAME)
PROJECT_SOURCE_DIR_DOCKER=/go/src/$(PROJECT_NAME)
PROJECT_BUILD_PATH=$(PROJECT_DIR)/build
PROJECT_BUILD_PATH_DOCKER=/go/build

#all 表示 make 默认会执行第一个目标
all: docker-build

build:
	$(GOBUILD) -ldflags "-X main.Version=$(TAG)" -o $(PROJECT_BUILD_PATH)/$(BINARY_NAME) $(BUILD_PATH)
test:
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(PROJECT_BUILD_PATH)/$(BINARY_NAME)
run:
	$(GOBUILD) -ldflags "-X main.Version=$(TAG)" -o $(PROJECT_BUILD_PATH)/$(BINARY_NAME) $(BUILD_PATH)
	$(PROJECT_BUILD_PATH)/$(BINARY_NAME)
deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop

# 使用 docker 进行程序的构建
docker-build:
	sudo docker run --rm -v $(PROJECT_SOURCE_DIR):$(PROJECT_SOURCE_DIR_DOCKER)	\
		-v $(PROJECT_BUILD_PATH):$(PROJECT_BUILD_PATH_DOCKER) $(GOLANG_VERSION)	\
		go build -ldflags "-X main.Version=$(TAG)" -o $(PROJECT_BUILD_PATH_DOCKER)/$(BINARY_NAME) $(BUILD_PATH)