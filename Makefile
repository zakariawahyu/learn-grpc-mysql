GO_WORKSPACE := $(GOPATH)/src

protoc:
	protoc --proto_path=model model/*.proto --go_out=$(GO_WORKSPACE) --go-grpc_out=$(GO_WORKSPACE)
	@echo "Protoc Compile Selesai"