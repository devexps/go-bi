# Not Support Windows

GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

SERVICE_RELATIVE_PATH=$(shell a=`basename $$PWD` && echo $$a)
SERVICE_DOCKER_IMAGE=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo "$$b/$$a:0.1.0")

API_PROTO_PATH=../api/proto
API_PROTO_FILES=$(shell cd $(API_PROTO_PATH) && find . -name *.proto)

# init env
init:
	@which micro &> /dev/null || go install github.com/devexps/go-micro/cmd/micro/v2@latest
	@which protoc-gen-go-http &> /dev/null || go install github.com/devexps/go-micro/cmd/protoc-gen-go-http/v2@latest

	@which protoc-gen-go &> /dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	@which protoc-gen-validate &> /dev/null || go install github.com/envoyproxy/protoc-gen-validate@v1.0.0
	@which protoc-gen-go-grpc &> /dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@which protoc-gen-openapi &> /dev/null || go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@which wire &> /dev/null || go install github.com/google/wire/cmd/wire@latest

# generate api proto
proto:
	@cd $(API_PROTO_PATH) && \
	mkdir -p ../gen ../gen/go ../gen/docs && \
	protoc --proto_path=./ \
		--proto_path=../third_party \
		--go_out=paths=source_relative:../gen/go \
		--go-http_out=paths=source_relative:../gen/go \
		--go-grpc_out=paths=source_relative:../gen/go \
		--go-errors_out=paths=source_relative:../gen/go \
		--openapi_out=fq_schema_naming=true,default_response=false:../gen/docs \
		$(API_PROTO_FILES)

# generate
generate:
	@go mod tidy
	@go generate ./...

# init and generate all
all:
	@make init
	@make proto
	@make generate

# run tests
test:
	@go test -race ./...

# run coverage tests
cover:
	@go test -v ./... -coverprofile=coverage.out

# build golang application
build:
	@go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

# clean build files
clean:
	@go clean
	@rm -f "coverage.out"

# build service app
app: proto generate build

# run debug application
run:
	@go run ./cmd/server -conf ./configs

# build docker image
docker:
	@docker build \
		-t $(SERVICE_DOCKER_IMAGE) \
    	-f ../.docker/Dockerfile \
    	--build-arg SERVICE_RELATIVE_PATH=$(SERVICE_RELATIVE_PATH) \
    	--build-arg GRPC_PORT=9090 \
    	--build-arg HTTP_PORT=8080 \
    	..

# run debug docker
docker-run:
	@docker run --rm -p 8080:8080 -p 9090:9090 \
		-v ${PWD}/../$(SERVICE_RELATIVE_PATH)/configs:/data/conf \
		$(SERVICE_DOCKER_IMAGE)

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help