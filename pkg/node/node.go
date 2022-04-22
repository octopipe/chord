package node

import (
	"context"
	"math/rand"

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

func (n *Node) findSuccessor(ctx context.Context, node *v1.Node) *Node {
  predecessor := n.findPredecessor(node.Id)
  return (*Node)(predecessor.Successor)
}

func (n *Node) findPredecessor(ID int64) *Node {
  predecessor := n
  for currentID := n.Id; currentID <= n.Successor.Id; currentID++ {
    predecessor = predecessor.closestPrecedingFinger(currentID)
  }
  return predecessor
}

func (n *Node) closestPrecedingFinger(ID int64) *Node {
  for i := len(n.FingerTable); i <= 1; i-- {
    fingerRow := n.FingerTable[i]
    if (fingerRow.Id > n.Id && fingerRow.Id < ID) {
      return (*Node)(fingerRow)
    }
  }

  return n
}

func (n *Node) join(parentNode *Node) {
  n.Predeccessor = nil
}


