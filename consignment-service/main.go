// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"sync"

// 	pb "github.com/alactic/shippydemo2/consignment-service//proto/consignment"
// 	vesselProto "github.com/alactic/shippydemo2/consignment-service/proto/vessel"
// 	"github.com/micro/go-micro"
// )

// const (
// 	port = ":50051"
// )

// type repository interface {
// 	Create(*pb.Consignment) (*pb.Consignment, error)
// 	GetAll() []*pb.Consignment
// }

// // Repository - Dummy repository, this simulates the use of a datastore
// // of some kind. We'll replace this with a real implementation later on.
// type Repository struct {
// 	mu           sync.RWMutex
// 	consignments []*pb.Consignment
// }

// // Create a new consignment
// func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
// 	repo.mu.Lock()
// 	updated := append(repo.consignments, consignment)
// 	repo.consignments = updated
// 	repo.mu.Unlock()
// 	return consignment, nil
// }

// // GetAll consignments
// func (repo *Repository) GetAll() []*pb.Consignment {
// 	return repo.consignments
// }

// // Service should implement all of the methods to satisfy the service
// // we defined in our protobuf definition. You can check the interface
// // in the generated code itself for the exact method signatures etc
// // to give you a better idea.
// type service struct {
// 	repo repository
// 	vesselClient vesselProto.VesselServiceClient
// }

// // CreateConsignment - we created just one method on our service,
// // which is a create method, which takes a context and a request as an
// // argument, these are handled by the gRPC server.
// func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

// 	// Here we call a client instance of our vessel service with our consignment weight,
// 	// and the amount of containers as the capacity value
// 	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
// 		MaxWeight: req.Weight,
// 		Capacity: int32(len(req.Containers)),
// 	})
// 	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
// 	if err != nil {
// 		return err
// 	}

// 	// We set the VesselId as the vessel we got back from our
// 	// vessel service
// 	req.VesselId = vesselResponse.Vessel.Id


// 	// Save our consignment
// 	consignment, err := s.repo.Create(req)
// 	if err != nil {
// 		return err
// 	}

// 	res.Created = true
// 	res.Consignment = consignment
// 	return nil
// }

// // GetConsignments -
// func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
// 	consignments := s.repo.GetAll()
// 	res.Consignments = consignments
// 	return nil
// }

// func main() {

// 	repo := &Repository{}

// 	// Set-up micro instance
// 	srv := micro.NewService(
// 		micro.Name("shippy.service.consignment"),
// 	)

// 	srv.Init()

// 	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", srv.Client())

// 	// Register handlers
// 	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

// 	// Run the server
// 	if err := srv.Run(); err != nil {
// 		fmt.Println(err)
// 	}
// }

// vessel-service/main.go
package main

import (
	// "context"
	// "errors"
	"fmt"

	// pb "github.com/alactic/shippydemo2/consignment-service/proto/vessel"
	// "github.com/micro/go-micro"
)

// // Repository interface creation
// type Repository interface {
// 	FindAvailable(*pb.Specification) (*pb.Vessel, error)
// }

// // VesselRepository interface creation
// type VesselRepository struct {
// 	vessels []*pb.Vessel
// }

// // FindAvailable - checks a specification against a map of vessels,
// // if capacity and max weight are below a vessels capacity and max weight,
// // then return that vessel.
// func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
// 	for _, vessel := range repo.vessels {
// 		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
// 			return vessel, nil
// 		}
// 	}
// 	return nil, errors.New("No vessel found by that spec")
// }

// // Our grpc service handler
// type service struct {
// 	repo repository
// }

// func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

// 	// Find the next available vessel
// 	vessel, err := s.repo.FindAvailable(req)
// 	if err != nil {
// 		return err
// 	}

// 	// Set the vessel as part of the response message type
// 	res.Vessel = vessel
// 	return nil
// }

// func main() {
// 	vessels := []*pb.Vessel{
// 		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
// 	}
// 	repo := &VesselRepository{vessels}

// 	srv := micro.NewService(
// 		micro.Name("shippy.service.vessel"),
// 	)

// 	srv.Init()

// 	// Register our implementation with 
// 	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

// 	if err := srv.Run(); err != nil {
// 		fmt.Println(err)
// 	}
// }

func main() {
	fmt.Println("Vessel application starting ...")
}