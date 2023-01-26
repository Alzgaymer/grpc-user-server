package main

import (
	"context"
	"net"
	"net/http"

	v1 "grpc-user-service/internal/controller/grpc/v1"

	protoUserService "github.com/Alzgaymer/grpc-rest-server/gen/go/proto/service/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcPortHost = "0.0.0.0:8082"
)

func main() {
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcPortHost)

	if err != nil {
		panic(err)
	}
	protoUserService.RegisterUserServiceServer(
		grpcServer,
		v1.NewUserServer(protoUserService.UnimplementedUserServiceServer{}),
	)

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = protoUserService.RegisterUserServiceHandlerFromEndpoint(context.Background(),
		mux,
		grpcPortHost,
		opts,
	)

	if err != nil {
		panic(err)
	}

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return grpcServer.Serve(l)
	},
	)
	g.Go(func() error {
		return http.ListenAndServe(":8081", mux)
	})

	err = g.Wait()
	if err != nil {
		panic(err)
	}
}
