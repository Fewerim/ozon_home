.PHONY: generate
generate:
	@mkdir -p product-service/pkg/api
	@protoc -I product-service/api \
	--go_out=product-service/pkg/api --go_opt=paths=source_relative \
    --go-grpc_out=product-service/pkg/api --go-grpc_opt=paths=source_relative \
    product-service.proto
