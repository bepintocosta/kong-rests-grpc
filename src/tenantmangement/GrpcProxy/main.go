package main

import (
	"context"
	"log"
	"net/http"
	"os"
	pb "sample/gen"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func CustomMatcher(key string) (string, bool) {

	switch key {
	case "X-Userinfo":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func httpResponseModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// set http status code
	if vals := md.HeaderMD.Get("x-http-code"); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, "x-http-code")
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
		w.WriteHeader(code)
	}

	return nil
}

func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {

	st := status.Convert(err)

	httpStatus := runtime.HTTPStatusFromCode(st.Code())

	w.WriteHeader(httpStatus)

	w.Write([]byte(st.Message()))
}

func main() {
	// Create a client connection to the gRPC server
	// This is where the gRPC-Gateway proxies the requests
	log.Println("run")

	conn, err := grpc.DialContext(
		context.Background(),
		os.Getenv("GRPC_HOST"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Println("Failed to dial server:", err)
	}

	log.Println("run conn")

	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption("application/json-patch+json", &runtime.JSONPb{
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: false, // explicit "false", &protojson.UnmarshalOptions{} would have the same effect
			},
		}),
		runtime.WithIncomingHeaderMatcher(CustomMatcher),
		runtime.WithForwardResponseOption(httpResponseModifier),
		runtime.WithErrorHandler(CustomErrorHandler),
	)
	log.Println("run gwmux")

	err = pb.RegisterTenantHandler(context.Background(), gwmux, conn)

	if err != nil {
		log.Fatalln("Failed to register tenant gateway:", err)
	}


	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
