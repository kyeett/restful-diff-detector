package main

import (
	"flag"
	"fmt"
	"github.com/kyeett/restful-diff-detector/grpc_client"
	"github.com/satori/go.uuid"
)

func main() {
	uuid := uuid.NewV4().String()

	idPtr := flag.String("id", uuid, "Client ID, used for logging")
	pathPtr := flag.String("path", "/user/1", "Path used for subscription")

	flag.Parse()

	fmt.Println(*idPtr)
	fmt.Println(*pathPtr)

	grpcclient.MakeFlow(*idPtr, *pathPtr)
}
