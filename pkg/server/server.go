package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/octopipe/dht/pkg/client"
	"github.com/octopipe/dht/pkg/node"
	chordPb "github.com/octopipe/dht/proto/chord/v1"
	v1 "github.com/octopipe/dht/proto/chord/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
  v1.ChordServer
  Node *node.Node
  Address string
  Client client.Client
}

func NewServer(node *node.Node, host string, port int) *Server {
  newServer := &Server{
    Address: fmt.Sprintf("%s:%d", host, port),
    Node: node,
    Client: client.NewClient(),
  }

  return newServer
}

func (s *Server) StartServer() error {
  lis, err := net.Listen("tcp", s.Address)
  if err != nil {
    return err
  }
  grpcServer := grpc.NewServer()
  chordPb.RegisterChordServer(grpcServer, s)
  return grpcServer.Serve(lis)
}

func (s *Server) Join(ctx context.Context, parentNodeAddress string) error {
  conn, err := s.Client.Connect(parentNodeAddress)
  if err != nil {
    return err
  }

  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()
  successor, err := conn.FindSuccessor(ctx, &v1.Node{Address: s.Node.Address, Id: s.Node.Id})
  if err != nil {
    return err
  }

  s.Node.Predeccessor = nil
  s.Node.Successor = successor
  return nil
}

func (s * Server) FindSuccessor(ctx context.Context, node *v1.Node) (*v1.Node, error) {
  return s.Node.FindSuccessor(node)
}


func (s * Server) Notify(ctx context.Context, node *v1.Node) (*emptypb.Empty, error) {
  s.Node.Notify(node)
  return nil, nil
}

