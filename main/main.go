package main

import (
	"encoding/json"
	"fmt"

	"github.com/rangertaha/hxe/internal/models"
	pb "github.com/rangertaha/hxe/internal/models/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Create a test service struct
	service := &models.Service{
		Name:        "test-service",
		Description: "A test service",
		// Command:     "echo",
		// Args:        "hello world",
		// Directory:   "/tmp",
		// User:        "testuser",
		// Group:       "testgroup",
		// Status:      models.ServiceStop,
		// Autostart:   true,
		// Enabled:     true,
		// MaxRetries:  3,
	}

	// Convert service struct to JSON
	jsonBytes, err := json.Marshal(service)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON:\n%s\n\n", string(jsonBytes))

	// Convert service struct to gRPC protobuf
	protoBytes, err := proto.Marshal(&pb.Service{service})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Protobuf:\n%x\n\n", protoBytes)

	// Unmarshal protobuf back to a new service struct
	unmarshalledService := &pb.Service{
		Name:        service.Name,
		Description: service.Description,
	}
	err = proto.Unmarshal(protoBytes, &unmarshalledService)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshalled protobuf:\n%+v\n\n", unmarshalledService)

	// Convert service struct to gRPC message
	grpcService := &pb.Service{
		Name:        service.Name,
		Description: service.Description,
	}

	fmt.Printf("gRPC message:\n%+v\n", grpcService)

	// Convert back from gRPC to service struct
	newService := &models.Service{
		Service: *grpcService,
	}

	fmt.Printf("\nConverted back to service struct:\n%+v\n", newService)
}
