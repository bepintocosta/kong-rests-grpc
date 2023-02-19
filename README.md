# kong-rests-grpc

In this repository, you can find the following steps to translate REST services to gRPC services using Kong.

# Cookbook

## HttpRules.yaml

The **HttpRules.yaml** file is used to define the routing between the REST API paths and the gRPC service methods.

Here is the basic structure of the **HttpRules.yaml file**:

```yaml 
type: google.api.Service
config_version: 3

http:
  rules:
    - selector: {package}.{service}.{method}
      get: /{path}
    - selector: {package}.{service}.{method}
      post: /{path}
      body: "*"
    - selector: {package}.{service}.{method}
      delete: /{path}
```
In the above example, you would replace **{package}**, **{service}**, and **{method}** with the actual names of your gRPC service and method. In the **get**, **post**, and **delete** lines, you would replace **{path}** with the appropriate URL path for your service.
If you choose the **post** method, you can use the **body** field to specify the expected structure of the request body. In the example, the **\*** character is used to indicate that any request body is acceptable. You can also specify a more detailed structure if needed.


## Generating the gRPC Proxy

Before you can generate the gRPC proxy, you need to install Go and Protoc on your machine. Here are the installation guides for both:

- GO: **https://go.dev/doc/install**
- Protoc: **https://grpc.io/docs/protoc-installation/**

Ensure you set these environment variables

```bash
root:~$ export GOPATH=/home/bcosta/go
root:~$ export GOBIN=$GOPATH/bin
root:~$ export PATH=$PATH:$GOPATH:$GOBIN
```

To install all dependencies needed to generate the rest client based in your **HttpRules.yaml**, you need to run these commands inside the  folder **"GrpcProxy/tools"** folder:

```bash
root:GrpcProxy/tools$ go mod init tools
root:GrpcProxy/tools$ go mod tidy
root:GrpcProxy/tools$ go generate tools.go
```

Next, go back to the "GrpcProxy" folder and run the following commands to generate the proxy:

```bash
root:GrpcProxy$ make genarate
```

The **make generate** command will execute the **protoc** command with the necessary plugins to generate the client stubs, as well as the reverse-proxy and OpenAPI spec for the REST API.

Now you can see in **"GrpcProxy/gen"** 4 files for every service inside your **HttpRules.yaml**

1. {service}_grpc.pb.go
2. {service}.pb.go
3. {service}.pb.gw.go
4. {service}.swagger.json

Now you need to call the registers handlers in **main.go** to do this, go to the generates **{service}.pb.gw.go** and find this method:

```go 
// RegisterAssetHandler registers the http handlers for service Asset to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func Register{YourService}Handler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAssetHandlerClient(ctx, mux, NewAssetClient(conn))
}
```

Now go to the **main.go** and replace the **pb.RegisterProviderHandler** to **pb.Register{YourService}Handler**
> ⚠️ **NOTE** You need to add a pb.Register for every service you have generate.

Now you can run the **make build** to produce an executable binary file and finally, the **make run** to run the application.
> ⚠️ **NOTE** You need to set the enviroment variable GRPC_HOST with url to your grpc server.

## Dockerfile

To create the Docker images for the gRPC service and gRPC proxy, you can use the dockerfiles provided in the project. Here's how to build the images:

1. Open a terminal window and navigate to the root directory of the project.
2. Execute the following commands, replacing **{yourtag}** with a unique tag for your images:

```bash
root:kong-rests-grpc$ docker  build -t {yourtag} -f src/{projectfolder}/GrpcAPI/Dockerfile .
root:kong-rests-grpc$ docker  build -t {yourtag} -f src/{projectfolder}/GrpcProxy/Dockerfile .
```

3. Wait for the images to be built. Once the build process is complete, you can use the **docker images** command to verify that the images were created successfully.

That's it! You now have Docker images for the gRPC service and gRPC proxy that you can use to deploy your application.

## Services in kubernetes

After building your Docker images, you can use the Kubernetes manifest files provided in the **GrpcAPI/k8s** folder to deploy your gRPC service and proxy. Here's how to make the necessary changes to the manifest files.

In the deployment kind, replace the following placeholders with your values:

- **{YourNamespace}**: The namespace you want to deploy the service to.
- **{PodName}**: The name of the pod that will be created.
- **{YourGrpcApiImage}**: The Docker image you built for the gRPC service.
- **{GrpcApiImageName}**: The name you want to give to the gRPC service container.
- **{YourGrpcProxyImage}**: The Docker image you built for the gRPC proxy.
- **{GrpcProxyImageName}**: The name you want to give to the gRPC proxy container.
- **{GrpcServerHost}**: The hostname of the gRPC server.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: {YourNamespace}
  creationTimestamp: null
  labels:
    app: {PodName}
  name: {PodName}
spec:
  replicas: 1
  revisionHistoryLimit: 3
  selector:
    matchLabels:
      app: {PodName}
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: {PodName}
    spec:
      containers:
      - image: "{YourGrpcApiImage}"
        imagePullPolicy: Always
        name: {GrpcApiImageName}
        resources: {}
        ports:
        - containerPort: 80
          name: http2
      - image: "{YourGrpcProxyImage}"
        imagePullPolicy: Always
        name: {GrpcProxyImageName}
        resources: {}
        ports:
        - containerPort: 8090
        env:
          - name:  GRPC_HOST
            value: "{GrpcServerHost}"
```
In the service kind, replace the following placeholders with your values:

- **{YourNamespace}**: The namespace you want to deploy the service to.
- **{ServiceName}**: The name you want to give to the service.
- **{PodName}**: The name of the pod that will be created.

```yaml
apiVersion: v1
kind: Service
metadata:
  namespace: {YourNamespace}
  name: {ServiceName}
  labels:
    run: {ServiceName}
    HealthChecks: "true"
spec:
  ports:
  - port: 8090
    name: proxy-port
    protocol: TCP
    targetPort: 8090
  - port: 5000
    name: grpc-port
    protocol: TCP
    targetPort: 80    
  selector:
    app: {PodName}
```

Once you've made the necessary changes, use the kubectl apply, command to apply the manifest files and deploy your gRPC service and proxy.

## Install Kong Ingress Controller


http://kong-admin-api.kong-admin.svc:8001/












