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
  if node.Id > n.Id || node.Id <= n.Successor.Id {
    return n.Successor, nil
  }

  closestNode := n.ClosestPrecedingFinger(node.Id)
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

func (n *Node) FindPredecessor(ID int64) *Node {
  predecessor := n
  for currentID := n.Id; currentID <= n.Successor.Id; currentID++ {
    predecessor = predecessor.ClosestPrecedingFinger(currentID)
  }
  return predecessor
}

func (n *Node) ClosestPrecedingFinger(ID int64) *Node {
  for i := len(n.FingerTable); i <= 1; i-- {
    fingerRow := n.FingerTable[i]
    if (fingerRow.Id > n.Id && fingerRow.Id < ID) {
      return (*Node)(fingerRow)
    }
  }

  return n
}

