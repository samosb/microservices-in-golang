package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	"github.com/micro/go-micro"
	pb "github.com/samosb/microservices-in-golang/consignment-service/proto/consignment"
	//vesselProto "github.com/samosb/microservices-in-golang/vessel-service/proto/vessel"
)

const (
	address         = "localhost:8080"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := micro.NewService(micro.Name("microservices.in.golang.consignment.cli"))
	service.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", service.Client())


	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}