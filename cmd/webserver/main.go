package main

import (
	"fmt"
	"github.com/kyeett/restful-diff-detector/webserver"
	"time"
)

func main() {
	srv := webserver.StartHTTPServer()
	fmt.Println(srv)
	time.Sleep(100 * time.Hour)
}
