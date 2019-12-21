package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	proto "github.com/vskut/twigo/pkg/grpc"
	"google.golang.org/grpc"
)

// RestGateway ...
type RestGateway struct{}

// NewRestGateway constructs the Server struct
func NewRestGateway() *RestGateway {
	return &RestGateway{}
}

// Run starts the rest gateway
func (gw *RestGateway) Run(endpoint, addr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	muxx := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := proto.RegisterAuthServiceHandlerFromEndpoint(ctx, muxx, endpoint, opts)
	if err != nil {
		return err
	}

	err = proto.RegisterTweetServiceHandlerFromEndpoint(ctx, muxx, endpoint, opts)
	if err != nil {
		return err
	}

	err = proto.RegisterUserServiceHandlerFromEndpoint(ctx, muxx, endpoint, opts)
	if err != nil {
		return err
	}

	if err := http.ListenAndServe(addr, muxx); err != nil {
		return err
	}

	return nil
}
