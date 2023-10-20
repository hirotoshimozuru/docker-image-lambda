GO_VERSION := $(shell cat .go-version)

go/vendor:
	@go mod vendor

build-aws: go/vendor
	docker build -t lambda-sample1 -f './sample1/Dockerfile' . --build-arg GO_VERSION=$(GO_VERSION) --target aws

build-local: go/vendor
	docker build -t lambda-sample1-local -f './sample1/Dockerfile' . --build-arg GO_VERSION=$(GO_VERSION) --target local

start-local:
	docker run --name lambda-sample1-local -p 9000:8080 --rm lambda-sample1-local /functions/main

stop-local:
	docker stop lambda-sample1-local

run-sample1:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations"


