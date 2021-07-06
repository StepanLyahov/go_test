.PHONY: openapi
openapi:
	oapi-codegen -generate types -o internal/adapter/delivery/http/v1/openapi_type.gen.go -package v1 api/openapi/getway.yaml
	oapi-codegen -generate chi-server -o internal/adapter/delivery/http/v1/openapi_server.gen.go -package v1 api/openapi/getway.yaml

protobuf:
	protoc -I api/protobuf calculator_uuid.proto --go_out=plugins=grpc:./internal/adapter/delivery/rpc