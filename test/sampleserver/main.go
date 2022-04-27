package main

import (
	"log"
	"os"
	"strconv"

	"github.com/octopipe/dht/pkg/node"
	"github.com/octopipe/dht/pkg/server"
)

func main() {
  node := node.NewNode()
  port, _ := strconv.Atoi(os.Getenv("PORT"))
  server := server.NewServer(node, os.Getenv("HOST"), port)

  if err := server.StartServer(); err != nil {
    log.Fatalln(err)
  }
}
