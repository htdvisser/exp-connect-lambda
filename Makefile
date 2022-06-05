.PHONY: default
default:

.PHONY: deps
deps:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
	go mod download

.PHONY: api
api:
	buf generate

.PHONY: build
build:
	mkdir -p dist
	go build -o dist/server ./cmd/connect-server
	GOOS=linux GOARCH=amd64 go build -o dist/lambda ./cmd/connect-lambda
	zip -j dist/lambda.zip dist/lambda

deploy/.terraform:
	terraform -chdir=deploy init

.PHONY: deploy
deploy: deploy/.terraform
	terraform -chdir=deploy fmt
	terraform -chdir=deploy validate
	terraform -chdir=deploy apply
