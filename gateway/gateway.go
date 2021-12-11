package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/leslesnoa/chatbox-grpc/pb"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "app:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterChatBoxHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	newMux := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"*"}),
	)(mux)
	return http.ListenAndServe(":8081", newMux)
}
func main() {
	flag.Parse()
	defer glog.Flush()
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
