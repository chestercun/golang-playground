package server

import (
	"context"
	"testing"

	pb "golang-playground/proto"
)

func TestGreeterService_SayHello(t *testing.T) {
	s := &GreeterService{}
	req := &pb.HelloRequest{Name: "World"}
	resp, err := s.SayHello(context.Background(), req)
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	if resp.Message != "Ahoy Matey! Hello, I see your name is World." {
		t.Errorf("Unexpected message: %s", resp.Message)
	}
}
