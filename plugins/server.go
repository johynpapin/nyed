package plugins

import (
	"context"
	pb "github.com/johynpapin/nyed/plugins/protobuf"
	"google.golang.org/grpc"
	"net"
)

type pluginServer struct {
	pb.UnimplementedPluginServer

	grpcServer *grpc.Server
}

func newPluginServer() *pluginServer {
	return &pluginServer{}
}

func (pluginServer *pluginServer) start() {
	listener, _ := net.Listen("tcp", ":4252")

	pluginServer.grpcServer = grpc.NewServer()
	pb.RegisterPluginServer(pluginServer.grpcServer, pluginServer)

	pluginServer.grpcServer.Serve(listener)
}

func (pluginServer *pluginServer) stop() {
	pluginServer.grpcServer.GracefulStop()
}

func (pluginServer *pluginServer) RegisterLanguage(ctx context.Context, request *pb.RegisterLanguageRequest) (*pb.RegisterLanguageReply, error) {
	return &pb.RegisterLanguageReply{}, nil
}
