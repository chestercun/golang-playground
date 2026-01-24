package server

import (
	"context"
	"fmt"

	pb "golang-playground/proto"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	message := "Ahoy Matey!"
	if req.Name != "" {
		message += fmt.Sprintf(" Hello, I see your name is %s.", req.Name)
	}
	if req.Email != "" {
		message += fmt.Sprintf(" And I see your email is %s.", req.Email)
	}
	if req.Age > 0 {
		message += fmt.Sprintf(" I also see that you are %d years old! Wow!", req.Age)
	}
	return &pb.HelloReply{Message: message}, nil
}
