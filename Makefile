# These are the values we want to pass for VERSION and BUILD
# git tag 1.0.1
# git commit -am "One more change after the tags"
VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`
BRANCH=`git branch --show-current`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD} -X main.Branch=${BRANCH}"

.PHONY : all
all : dep build

version:
	@git describe --tags --abbrev=0

git:
	@git add .
	@git commit -m "$m"
	@git push -u origin ${BRANCH}

debug :
	go run ${LDFLAGS} cmd/service/main.go -debug

debug-grpc-cli :
	go run ${LDFLAGS} cmd/client-grpc/main.go -debug

lint: check-lint dep
	golangci-lint run --timeout=5m -c .golangci.yml

build:
	@go build ${LDFLAGS} -o service ./cmd/service
	@build ${LDFLAGS} -o grpc-client ./cmd/ client-grpc
	@build ${LDFLAGS} -o http-client ./cmd/client-http

dep:
	@go mod tidy -compat=1.17
	@go mod download

test:
	@go test -v -timeout 30s ./...

coverage:
	@go test -timeout 30s ./... -covermode=atomic

check-lint:
	@which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

proto:
	@protoc --proto_path=api/grpc/v1 --proto_path=third_party  --go_out=. --go-grpc_out=. --grpc-gateway_out=. service.proto
	@protoc --proto_path=api/grpc/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 service.proto
