APP=vvvstore
SHELL := /bin/bash
BASEDIR = $(shell pwd)

.PHONY: build run clean docs install help

all: build

build: clean
	@echo $(shell go version)
	@go build -v -tags=jsoniter ./cmd/${APP}

run:
	@go run ./cmd/${APP} -x=true -m=debug

clean:
	rm -rf ./${APP} ./${APP}.exe
	@go clean -x

docs:
	@swag init

install:
	@go env -w GO111MODULE=on
	@go env -w GOPROXY=https://goproxy.cn,direct
	@go mod download

help:
	@echo "make 编译包和依赖项"
	@echo "make run 直接运行程序"
	@echo "make clean 清除编译文件和缓存文件"
	@echo "make docs 生成Swagger依赖文件"
	@echo "make install 下载并安装依赖包"