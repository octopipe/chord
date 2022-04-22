package main

import (
	"flag"
	"log"

	"github.com/octopipe/chord/pkg/node"
	"github.com/octopipe/chord/pkg/server"
)

var (
  address string
)

func init() {
  flag.StringVar(&address, "address", "", "Server address")
  flag.Parse()
}

func main() {
  server  := server.NewServer(address)
  node := node.NewNode(server)

  if err := node.StartServer(); err != nil {
    log.Fatalln(err)
  }
}


