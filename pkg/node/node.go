package node

import (
	"context"
	"math/rand"

	"github.com/octopipe/chord/pkg/client"
	v1 "github.com/octopipe/chord/proto/chord/v1"
)

type Node v1.Node

func NewNode() Node {
  return Node {
  }
}

func newId() int64 {
  return rand.Int63n(100)
}

func (n *Node) FindSuccessor(node *v1.Node) (*v1.Node, error) {
  if Between(node.Id, n.Id, n.Successor.Id) {
    return n.Successor, nil
  }

  closestNode := n.ClosestPrecedingNode(node.Id)
  closestNodeServer, err := client.NewClient().Connect(closestNode.Address)
  if err != nil {
    return nil, err
  }

  successor, err := closestNodeServer.FindSuccessor(context.Background(), node)
  if err != nil {
    return nil, err
  }

  return successor, nil
}

func (n *Node) ClosestPrecedingNode(id int64) *Node {
  m := len(n.FingerTable)

  for i := m; i >= 0; i-- {
    row := n.FingerTable[i]

    if Between(row.Id, n.Id, id) {
      return (*Node)(row)
    }
  }

  return n
}

func Between(id, start, end int64) bool {
  return id > start || id <= end
}

