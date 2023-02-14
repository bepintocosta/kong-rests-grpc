# kong-rests-grpc

In this repository, you can find the following steps to translate REST services to gRPC services using Kong.

# Cookbook

1. After creating your grpc services you need to create your HttpRules.yaml.
The best practice is to create a folder on the same level as yours grpc services.

So in this repo I create a folder GrpcProxy and inside that folder, I put the HttpRules file.

In httpRules file pastes this yaml.

```yaml 
type: google.api.Service
config_version: 3

http:
  rules:
    - selector: AssetManagement.Asset.ListAssets
      get : /{tenantCode}/asset
    - selector: AssetManagement.Asset.CreateAsset
      post : /{tenantCode}/asset
      body : "*"
    - selector: AssetManagement.Provider.CreateProvider
      post : /provider
      body: "*"
    - selector: AssetManagement.Provider.ListProviders
      get : /provider
```

The list of rules needs to be:
Selector - This is the method of your service and you need to put this structure {package}.{service}.{method}
Http Method -  In the second line you need to put the HTTP method your service responds to. (get, post, delete)
In case you choose the post method, you need to put body, in our scenario you always put * to the mapping of all the fields.

Now you need to generate our proxy.

First of all, install go and protoc in your machine

GO : https://go.dev/doc/install
Protoc: https://grpc.io/docs/protoc-installation/

After that create this folder GrpcProxy/tools and create a file tool.go and put this code:


```go

//go:generate go install github.com/bufbuild/buf/cmd/buf github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc

package tools

import (
    "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
    "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
    "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
    "google.golang.org/protobuf/cmd/protoc-gen-go"
	"github.com/bufbuild/buf/cmd/buf"
)
```

After this run this commands

```bash

root:GrpcProxy/tools$ go mod init tools
root:GrpcProxy/tools$ go mod tidy
go generate tools.go
```

Now run 

```bash
root:GrpcProxy$ make genarate
root:GrpcProxy$ make build
root:GrpcProxy$ make run
```

The first command will execute the protoc command and create the stubs, the others two will be build and run the application.

Note: In make file change the variable GRPC_HOST to our grpc server url.













