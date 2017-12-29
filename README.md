# restful-diff-detector
An application that polls any REST interface and allows other applications to subscribe to changes to the monitored interface and receive notifications over gRPC.

**Example sequence**

![alt text](sequence.png "Example of sequence")

1. **REST interface** - Any REST interface accessible over HTTP
2. **restful-diff-detector**
  - Polls **REST interface** at address http://www.yourrestapi.com/
- **clients**
  - Connect to **restful-diff-detector** using gRPC
  - Subscribe to a path in the REST API, for example "/1.1/tweets/0"
  - Receives updates over gRPC from **restful-diff-detector** if subscribed path is updated

# Features
1. Polling of REST interface
2. Diff result from REST interface
3. Serve gRPC interface to clients
5. Create example client in Go
4. Create subscription for path in REST

# Lessons learnt (not related to application)
### Install
go get -u github.com/golang/lint/golint

### Update sequence diagram
java -jar plantuml.jar sequence.puml

### References
[5 tricks for tests in golang](https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742)

[pub-sub testing](https://github.com/cskr/pubsub/blob/master/pubsub_test.go)


[golang tutorial](https://tour.golang.org/)
