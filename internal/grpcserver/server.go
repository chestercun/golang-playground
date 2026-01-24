package grpcserver

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"golang-playground/server"
	pb "golang-playground/proto"
)

const bufSize = 1024 * 1024

type InProcessGRPC struct {
	Listener *bufconn.Listener
	Server   *grpc.Server
}

func New() *InProcessGRPC {
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server.GreeterService{})

	return &InProcessGRPC{
		Listener: lis,
		Server:   s,
	}
}

func (g *InProcessGRPC) Start() {
	go func() {
		_ = g.Server.Serve(g.Listener)
	}()
}

func (g *InProcessGRPC) DialContext(ctx context.Context) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		ctx,
		"in-process",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return g.Listener.Dial()
		}),
		grpc.WithInsecure(), // acceptable for in-process usage
	)
}
