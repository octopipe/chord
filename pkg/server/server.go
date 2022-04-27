package server

import (
	"context"
	"fmt"
	"log"
	"net"

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

type ServerConfig struct {
  Node *node.Node
  IsRoot bool
  Host string
  Port int
  ParentNodeAddress string
}

func NewServer(config ServerConfig) (*Server, error) {
  newServer := &Server{
    Address: fmt.Sprintf("%s:%d", config.Host, config.Port),
    Node: config.Node,
    Client: client.NewClient(),
  }

  if config.IsRoot {
    newServer.Create()
  } else {
    err := newServer.Join(context.Background(), config.ParentNodeAddress)
    if err != nil {
      return nil, err
    }
  }

  return newServer, nil
}

func (s *Server) StartServer() error {
  log.Println(fmt.Sprintf("Starting server on %s", s.Address))
  lis, err := net.Listen("tcp", s.Address)
  if err != nil {
    return err
  }
  grpcServer := grpc.NewServer()
  chordPb.RegisterChordServer(grpcServer, s)
  return grpcServer.Serve(lis)
}

func (s *Server) Create() {
  log.Println("Create network")
  s.Node.Predeccessor = nil
  s.Node.Successor = (*v1.Node)(s.Node)
}

func (s *Server) Join(ctx context.Context, parentNodeAddress string) error {
  log.Println(fmt.Sprintf("Join to network by %s", parentNodeAddress))
  conn, err := s.Client.Connect(parentNodeAddress)
  if err != nil {
    return err
  }

  successor, err := conn.FindSuccessor(ctx, &v1.Node{Address: s.Address, Id: s.Node.Id})
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

