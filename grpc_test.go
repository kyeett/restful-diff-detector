package main

import (
	"context"
	"fmt"
	"github.com/kyeett/restful-diff-detector/grpc_server"
	pb "github.com/kyeett/restful-diff-detector/proto"
	"testing"
)

func TestHello(t *testing.T) {
	s := grpcserver.Server{}

	// set up test cases
	tests := []struct {
		name string
		want string
	}{
		{
			name: "world",
			want: "Hello world",
		},
		{
			name: "123",
			want: "Hello 123",
		},
	}

	for _, tt := range tests {
		req := &pb.HelloRequest{Name: tt.name}
		resp, err := s.SayHello(context.Background(), req)
		if err != nil {
			t.Error("HelloTest() got unexpected error")
		}
		if resp.Message != tt.want {
			t.Errorf("HelloText(%v)=%v, wanted %v", tt.name, resp.Message, tt.want)
		}
	}
}

func TestSubscribe(t *testing.T) {
	s := grpcserver.Server{}

	req := &pb.DiffSubscribe{Path: "/user/1"}
	resp, err := s.Subscribe(context.Background(), req)
	if err != nil {
		t.Error("HelloTest() got unexpected error")
	}

	fmt.Printf("%v\n", resp.Message)
}
