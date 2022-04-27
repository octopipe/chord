package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/octopipe/dht/pkg/node"
	"github.com/octopipe/dht/pkg/server"
)

var (
  host = ""
  port = ""
  parent = ""
)

func init() {
  flag.StringVar(&host, "host", "", "Host")
  flag.StringVar(&port, "port", "", "Port")
  flag.StringVar(&parent, "parent", "", "Node in chord network")
  flag.Parse()
}

func main() {
  node := node.NewNode()
  portParsed, _ := strconv.Atoi(port)
  serverConfig := server.ServerConfig{
    Node: node,
    Host: host,
    Port: portParsed,
    ParentNodeAddress: parent,
    IsRoot: parent == "",
  }
  server, err := server.NewServer(serverConfig)
  if err != nil {
    log.Fatalln(err)
  }

  if err := server.StartServer(); err != nil {
    log.Fatalln(err)
  }
}
