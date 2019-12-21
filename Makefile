build:
	@mkdir -p bin
	@env go build -o bin ./cmd/...

build/linux:
	@mkdir -p bin
	@env GOOS=linux go build -o bin ./cmd/...

dockerize:
	@docker-compose -f deployments/docker-compose.yml up

docker/build:
	@docker-compose -f deployments/docker-compose.yml build

run/server:
	@go run ./cmd/grpc-server/main.go

run/client:
	@go run ./cmd/rest-client/main.go

run/swagger:
	swagger serve ./api/swagger-spec/api.json

proto:
	protoc -I . ./pkg/grpc/proto/*.proto \
	 -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
	 -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --go_out=plugins=grpc:. \
	 --validate_out="lang=go:." \
	 --grpc-gateway_out=logtostderr=true:. \
	 --swagger_out=logtostderr=true:.
	swagger mixin ./pkg/grpc/proto/*.swagger.json --output=./api/swagger-spec/api.json

test:
	go test ./... -cover

test/coverage:
	mkdir -p tmp
	go test -v ./... -coverprofile=./tmp/cover.out
	go tool cover -html=./tmp/cover.out -o ./tmp/coverout.html
	open ./tmp/coverout.html
	@echo "\n"

test/clear:
	rm -rf ./tmp


.PHONY: run run/server run/client proto swagger test test/coverage test/clear build build/linux dockerize docker/build