module github.com/samosb/microservices-in-golang/consignment-service

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.10.0
	github.com/nats-io/nats-server/v2 v2.1.0 // indirect
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	google.golang.org/grpc v1.23.1 // indirect
)

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
