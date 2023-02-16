generate:
	protoc -I ./protos \
	--go_out ./gen \
	--go_opt paths=source_relative \
	--go-grpc_out ./gen \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./gen \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt grpc_api_configuration=./httpRules.yaml \
	  ./protos/*.proto \
	--openapiv2_out ./gen \
    --openapiv2_opt logtostderr=true \
    --openapiv2_opt use_go_templates=true 

generate-local:
	protoc -I ../../../protos \
	--go_out ./gen \
	--go_opt paths=source_relative \
	--go-grpc_out ./gen \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./gen \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt grpc_api_configuration=./httpRules.yaml \
	  ../../..//protos/*.proto \
	--openapiv2_out ./gen \
    --openapiv2_opt logtostderr=true \
    --openapiv2_opt use_go_templates=true 

clean:
	 rm gen/*.go

build: 
	go build main.go
	
run: 
	go run main.go

client: 
	go build -o grpc_gateway main.go

