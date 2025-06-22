package exec

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"

	pb "github.com/rangertaha/hxe/internal/plugins/builtins/exec/proto"
)

// Service is the interface that we're exposing as a plugin.
type Service interface {
	Init(args *structpb.Struct) error
	Start() error
	Stop() error
	Restart() error
	Enable() error
	Disable() error
	Stdin(data []byte) error
	Stdout() ([]byte, error)
}

// ServicePlugin is the implementation of plugin.GRPCPlugin
type ServicePlugin struct {
	plugin.Plugin
	Impl Service
}

func (p *ServicePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterServiceServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *ServicePlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: pb.NewServiceClient(c)}, nil
}

// Handshake is a common handshake that is shared between plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SERVICE_PLUGIN",
	MagicCookieValue: "hxe",
}

// GRPCServer is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	pb.UnimplementedServiceServer
	Impl Service
}

func (m *GRPCServer) Init(ctx context.Context, req *pb.InitRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, m.Impl.Init(req.Args)
}

func (m *GRPCServer) Start(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, m.Impl.Start()
}

func (m *GRPCServer) Stop(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, m.Impl.Stop()
}

func (m *GRPCServer) Restart(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, m.Impl.Restart()
}

func (m *GRPCServer) Enable(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, m.Impl.Enable()
}

func (m *GRPCServer) Disable(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, m.Impl.Disable()
}

func (m *GRPCServer) Stdin(ctx context.Context, req *pb.StdinRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, m.Impl.Stdin(req.Data)
}

func (m *GRPCServer) Stdout(ctx context.Context, req *emptypb.Empty) (*pb.StdoutResponse, error) {
	data, err := m.Impl.Stdout()
	if err != nil {
		return nil, err
	}
	return &pb.StdoutResponse{Data: data}, nil
}

// GRPCClient is an implementation of Service that talks over gRPC.
type GRPCClient struct {
	client pb.ServiceClient
}

func (m *GRPCClient) Init(args *structpb.Struct) error {
	_, err := m.client.Init(context.Background(), &pb.InitRequest{Args: args})
	return err
}

func (m *GRPCClient) Start() error {
	_, err := m.client.Start(context.Background(), &emptypb.Empty{})
	return err
}
func (m *GRPCClient) Stop() error {
	_, err := m.client.Stop(context.Background(), &emptypb.Empty{})
	return err
}

func (m *GRPCClient) Restart() error {
	_, err := m.client.Restart(context.Background(), &emptypb.Empty{})
	return err
}

func (m *GRPCClient) Enable() error {
	_, err := m.client.Enable(context.Background(), &emptypb.Empty{})
	return err
}

func (m *GRPCClient) Disable() error {
	_, err := m.client.Disable(context.Background(), &emptypb.Empty{})
	return err
}

func (m *GRPCClient) Stdin(data []byte) error {
	_, err := m.client.Stdin(context.Background(), &pb.StdinRequest{Data: data})
	return err
}

func (m *GRPCClient) Stdout() ([]byte, error) {
	resp, err := m.client.Stdout(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
