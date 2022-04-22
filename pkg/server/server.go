package server

import (
	"net"

	"github.com/octopipe/chord/pkg/node"
	chordPb "github.com/octopipe/chord/proto/chord/v1"
	"google.golang.org/grpc"
)

type Server struct {
  node.Node
  Address string
}

func NewServer(address string) Server {
  newServer := Server{
    Address: address,
  }

  return newServer
}

func (s Server) StartServer() error {
  lis, err := net.Listen("tcp", s.Address)
  if err != nil {
    return err
  }
  grpcServer := grpc.NewServer()
  chordPb.RegisterChordServer(grpcServer, s)
  return grpcServer.Serve(lis)
}

