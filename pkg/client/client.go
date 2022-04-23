package client

import (
	chordPb "github.com/octopipe/chord/proto/chord/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
  credentials credentials.TransportCredentials
}

func NewClient() Client {
  return Client {
    credentials: insecure.NewCredentials(),
  }
}

func (c Client) Connect(address string) (chordPb.ChordClient, error) {
  conn, err := grpc.Dial(address, grpc.WithTransportCredentials(c.credentials))
  if err != nil {
    return nil, err
  }

  return chordPb.NewChordClient(conn), nil
}
