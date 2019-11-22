package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/alactic/shippydemo2/consignment-service/proto/consignment"
	vesselProto "github.com/alactic/shippydemo2/vessel-service/proto/vessel"
	"github.com/alactic/shippydemo2/consignment-service/datastore"
	"github.com/micro/go-micro"
)

const (
	port = ":50051"
	defaultHost = "datastore:27017"
)

func main() {

	// repo := &Repository{}

	// Set-up micro instance
	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := datastore.CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}

	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", srv.Client())

	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Register handlers
	// pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}