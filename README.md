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

### Custom Resource Definition
To install the kong ingress controller you need to create some custom resource definitions in your Kubernetes cluster.
In root you have the **k8s** folder inside this folder, which exists a bash file with the name **k8s_apply_crds.sh**, you must run this file to apply the crd's to your k8s cluster.

```bash
root:kong-rests-grpc/k8s$ ./k8s_apply_crds.sh
```
### Create Gateways
Once you have completed the previous steps, the next step is to apply the Kong gateway to your Kubernetes (k8s) cluster. In this example, we will create two Gateways and install them in separate namespaces. To accomplish this, we will use sed commands to replace all the tokens inside the template files with respective folders.
In the template file, you will notice the token **{{KONG_CONTEXT}}**. The sed command will replace this token with either "public" or "admin", depending on the namespace.
After the token replacement, you can find the modified files in the respective "admin" and "public" folders.

In the folders, you will find two shell scripts named **k8s_apply_admin.sh** and **k8s_apply_public.sh**. After the token replacement, you can run these scripts to create the gateways in your Kubernetes cluster, but let me explain some detail inside this files.

The following commands are present inside the shell script:

```bash
kubectl -n kong-admin create configmap kong-plugins-jwt-keycloak --from-file=../plugins/jwt-keycloak --dry-run=client -o yaml | kubectl apply -f -
kubectl -n kong-admin create configmap kong-plugins-jwt-keycloak-validators --from-file=../plugins/jwt-keycloak/validators --dry-run=client -o yaml | kubectl apply -f -
```

These commands are used to create a config map with the necessary files for installing plugins in the Kong gateway. In this example, the installed plugin is **jwt-keycloak**, and the required files are present in the **jwt-keycloak** folder inside the **plugins** folder.

You need to pass these files inside the container where the gateway will be installed, so first, you need to create a config map. Now, it is necessary to create a volume with this information, and the relationship is created in the **Deployment.yaml** file.

The following YAML code snippets demonstrate how to do this.

```yaml
volumes:
      ...
      - name: kong-plugins-jwt-keycloak
        configMap:
          name: kong-plugins-jwt-keycloak
      - name: kong-plugins-jwt-keycloak-validators
        configMap:
          name: kong-plugins-jwt-keycloak-validators
```

Next, you need to create a volume mount in the **ingress-kong** container.

```yaml
volumeMounts:
        - name: kong-plugins-jwt-keycloak
          mountPath: /opt/kong/plugins/jwt-keycloak
        - name: kong-plugins-jwt-keycloak-validators
          mountPath: /opt/kong/plugins/jwt-keycloak/validators
```

Additionally, you need to enable the plugin by setting the following environment variables in the same location.

```yaml
      - name: KONG_PLUGINS
        value: bundled,jwt-keycloak
      - name: KONG_LUA_PACKAGE_PATH
        value: /opt/?.lua;;
```
Note that the **KONG_LUA_PACKAGE_PATH** should always have this value, but in the **KONG_PLUGINS** variable, you can add more plugins such as bundled, jwt-keycloak, otherplugin.

To apply this plugin to a Kong gateway, you can use the **KongClusterplugin-jwt-keycloak** as a template. Note that this plugin only applies to a **kong-public gateway**. You can find the relevant script to apply this plugin in the public/k8s_apply_public.sh file.

### Ingress Rules

To create ingress rules inside the kong ingress controller, you can use the template file located in the **k8s** folder under the **GrpcApi** folder of **tenantmangement**.

This template file includes the creation of external service within the **kong-admin** namespace.


```yaml
kind: Service
apiVersion: v1
metadata:
  name: tenant-management-grpc-api-proxy
  namespace: kong-admin
  annotations:
    "konghq.com/override": "mission-health-checking"
spec:
  type: ExternalName
  externalName: tenant-management-grpc-api-service.netponto.svc
  ports:
     - port: 8090
       targetPort: 8090
       protocol: TCP 
```

This creates an external service named **tenant-management-grpc-api-proxy** in the **kong-admin** namespace, with a target port of 8090. Note that the konghq.com/override annotation is used to override the default health checking behavior of Kong. You can modify this file to suit your specific use case.

After the service is created, the next step is to create a Kong Ingress that will provide health-checking for the external service. The following YAML file creates a KongIngress object with a health-check configuration:

```yaml
apiVersion: configuration.konghq.com/v1
kind: KongIngress
metadata:
    name: health-checking
    namespace: kong-admin
upstream:
  healthchecks:
    passive:
      healthy:
        successes: 3
        http_statuses: [ 200, 201, 203, 204, 400, 401, 404 ]
      unhealthy:
        http_failures: 3
        http_statuses: [ 500, 502, 503 ]
        tcp_failures: 1
        timeouts: 1
```

This health-check configuration specifies that the service will be considered healthy after three successful requests with any of the HTTP statuses in the **healthy** field. Conversely, if the service receives three requests with any of the HTTP statuses in the **unhealthy** field, it will be considered unhealthy. The **tcp_failures** and **timeouts** fields specify the number of TCP failures and timeouts that will cause the service to be considered unhealthy.

The next section of the template defines an Ingress rule:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: tenant-post
  namespace: kong-admin
  annotations:
    kubernetes.io/ingress.class: kong-admin
spec:
  rules:
  - http:
      paths:
      - path: /([A-Z, a-z, 0-9]{1,4})/tenant
        pathType: Prefix
        backend:
          service:
             name: tenant-management-grpc-api-proxy
             port:
               number: 8090
```

This rule specifies that all requests with the path **/([A-Z, a-z, 0-9]{1,4})/tenant** should be routed to the **tenant-management-grpc-api-proxy** service on port **8090**. The **pathType** is set to **Prefix** which means that any path that starts with **/([A-Z, a-z, 0-9]{1,4})/tenant** will be matched by this rule. The **kubernetes.io/ingress.class** annotation is used to specify that this Ingress rule should be handled by the Kong Ingress controller. The **namespace** field specifies that this Ingress rule should be created in the **kong-admin** namespace.

In next section of the yaml:

```yaml
apiVersion: configuration.konghq.com/v1
kind: KongIngress
metadata:
  name: tenant-post
  namespace: kong-admin
route:
  methods:
  - POST
  - OPTIONS
  regex_priority: 0
```

Specifies a KongIngress, that only the HTTP methods **POST** and **OPTIONS** are allowed for the path **/([A-Z, a-z, 0-9]{1,4})/tenant**. It also sets the **regex_priority** to 0, which determines the order in which regular expressions are evaluated by Kong when multiple routes match a request.

Now you only need to execute the following command

```bash
root:kong-rests-grpc/src/tenantmangement/GrpcAPI/k8s/grpc-http-rules.yml$ kubectl apply -f grpc-http-rules.yml
```



http://kong-admin-api.kong-admin.svc:8001/












