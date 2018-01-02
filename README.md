# restful-diff-detector
An application that polls any REST interface and allows other applications to subscribe to changes to the monitored interface and receive notifications over gRPC.

## Description
- **REST interface** - Any REST interface accessible over HTTP
- **restful-diff-detector**
  - Polls **REST interface** at address http://www.yourrestapi.com/
- **clients**
  - Connect to **restful-diff-detector** using gRPC
  - Subscribe to a path in the REST API, for example "http://www.yourrestapi.com//1.1/tweets/0"
  - Receives updates over gRPC from **restful-diff-detector** if subscribed path is updated

**Example sequence**

![alt text](sequence.png "Example of sequence")

# Features
### ~~Release 0.1~~
1. ~~Update webserver through web call~
4. ~~Create subscription for path in REST~~
1. ~~Polling of REST interface~~
5. ~~Create example client in Go~~
2. ~~Diff result from REST interface~~
1. ~~Serve gRPC interface to clients~~

### Release 0.2
1. Use `dep`
5. Use `multi-stage docker build`
1. Make `make container` build work
2. Make `make push` work
3. Make `make deploy` work
1. IP addr as input to webserver
6. Run webserver in docker

### Release 0.3
1. Use some kind of versioning system 
1. Port as input to client
3. Port as input to server
2. Run tests in travis-ci
4. Update /user/ after X accesses
1. Handle HTTP server not available (server)

### Release 0.x
1. Handle HTTP crashing (server)
2. Handle path doesn't exist (server)
3. Send errors to client

### Unplanned
6. Debian package per application
4. Use this lib for generating JSON patches instead of changes. http://jsonpatch.com/

# Lessons learnt (not related to application)

### Do something every X seconds
```
ticker := time.NewTicker(1 * time.Second)
for range ticker.C {
    // Do stuff
}
```

### Install
go get -u github.com/golang/lint/golint

### Install Python
```
pip install grpcio-tools
python -m grpc_tools.protoc -I proto --python_out=proto  --grpc_python_out=proto proto/hello.proto
```

### Compile protobuf
```
protoc --go_out=. proto/hello.proto
```

### Update sequence diagram
```
java -jar plantuml.jar sequence.puml
```


### Project structure

[Respository structure](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)

[Go build template](https://github.com/thockin/go-build-template)

### References
[5 tricks for tests in golang](https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742)

[pub-sub testing](https://github.com/cskr/pubsub/blob/master/pubsub_test.go)

[golang tutorial](https://tour.golang.org/)

[go: how to shutdown http server](https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve)

[testing a gPRC service](https://stackoverflow.com/questions/42102496/testing-a-grpc-service)

[gRPC error handling](http://avi.im/grpc-errors)

[Setting backoff parameters in gPRC](https://github.com/grpc/grpc/issues/11277)


