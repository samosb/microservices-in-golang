module github.com/samosb/microservices-in-golang

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.10.0
	github.com/micro/protobuf v0.0.0-20180321161605-ebd3be6d4fdb // indirect
	github.com/micro/protoc-gen-micro v0.8.0 // indirect
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	google.golang.org/grpc v1.23.1
)

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
