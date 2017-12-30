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

package main

import (
	"fmt"
	pb "github.com/kyeett/restful-diff-detector/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func makeFlow() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDiffSubscriberClient(conn)
	req := &pb.DiffSubscribe{Path: "/user/1"}
	stream, err := client.SubscribeStream(context.Background(), req)

	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
	}

}

func makeCall(ch chan string) {
	start := time.Now()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	ch <- fmt.Sprintf("Greeting: %s\t(took %d ms)", r.Message, int64(time.Since(start)/time.Millisecond))
}

func clientMain() {
	start := time.Now()
	ch := make(chan string)

	numCalls := 10
	for n := 0; n < numCalls; n++ {
		go makeCall(ch)
	}
	fmt.Println("All calls sent")

	for n := 0; n < numCalls; n++ {
		fmt.Printf("%v\n", <-ch)
	}
	fmt.Printf("Got all greetings in %d ms", int64(time.Since(start)/time.Millisecond))
}

func main() {
	makeFlow()
	//clientMain()
}
