package grpc

import (
	"google.golang.org/grpc"
	"net"
	"signal-server/internal/services"
)

type Server struct {
	grpcServer    *grpc.Server
	signalService *services.SignalService
}

func NewServer(signalService *services.SignalService) *Server {
	grpcServer := grpc.NewServer()
	signalServer.RegisterSignalServer(grpcServer, signalService)
	return &Server{grpcServer: grpcServer, signalService: signalService}
}

func (s *Server) Start(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(lis)
}
