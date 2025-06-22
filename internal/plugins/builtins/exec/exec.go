package exec

import (
	"log"

	// "github.com/hashicorp/go-plugin"
	"google.golang.org/protobuf/types/known/structpb"
)

// ServiceImpl is the actual implementation of our service.
type ServiceImpl struct{}

func (s *ServiceImpl) Init(args *structpb.Struct) error {
	log.Println("Service plugin initialized with args:", args)
	return nil
}

func (s *ServiceImpl) Stdin() error {
	log.Println("Service started")
	return nil
}

func (s *ServiceImpl) Stdcmd() error {
	log.Println("Service stopped")
	return nil
}

func (s *ServiceImpl) Stdout() error {
	log.Println("Service stdout")
	return nil
}

func (s *ServiceImpl) Stderr() error {
	log.Println("Service stderr")
	return nil
}

// func main() {
// 	plugin.Serve(&plugin.ServeConfig{
// 		HandshakeConfig: Handshake,
// 		Plugins: map[string]plugin.Plugin{
// 			"service": &ServicePlugin{Impl: &ServiceImpl{}},
// 		},
// 		GRPCServer: plugin.DefaultGRPCServer,
// 	})
// }
