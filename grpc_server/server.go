/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpcserver

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "github.com/kyeett/restful-diff-detector/proto"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// Server is used to implement hello.GreeterServer.
type Server struct{}

// SayHello implements hello.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("Received request from %v\n", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// SayHello implements hello.GreeterServer
func (s *Server) Subscribe(ctx context.Context, in *pb.DiffSubscribe) (*pb.DiffNotification, error) {
	fmt.Printf("Received request from %v\n", in.Path)
	return &pb.DiffNotification{ResponseData: "Hello hello, " + in.Path}, nil
}

// ListFeatures lists all features contained within the given bounding Rectangle.
func (s *Server) SubscribeStream(in *pb.DiffSubscribe, stream pb.DiffSubscriber_SubscribeStreamServer) error {

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {

		if err := stream.Send(&pb.DiffNotification{ResponseData: "Hello hello, " + in.SubscriberId}); err != nil {
			return err
		}
		fmt.Println("Sending message to ", in.SubscriberId, "with information about", in.Path)
		time.Sleep(200 * time.Millisecond)

	}
	return nil
}

func ServerMain() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterGreeterServer(s, &Server{})
	pb.RegisterDiffSubscriberServer(s, &Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	ServerMain()
}
