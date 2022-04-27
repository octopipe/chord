package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
)

type ServerContainer struct{
  testcontainers.Container
  URI string
}

func setupServer(ctx context.Context, port int, parent string) (*ServerContainer, error) {
  parsedPort := fmt.Sprintf("%d/tcp", port)
  dockerfileReq := testcontainers.FromDockerfile {
    Context: "..",
    Dockerfile: "test/sampleserver/Dockerfile",
  }
  req := testcontainers.ContainerRequest {
    FromDockerfile: dockerfileReq,
    ExposedPorts: []string{parsedPort},
    Env: map[string]string{
      "PORT": parsedPort,
      "HOST": "0.0.0.0",
      "PARENT": parent,
    },
  }
  container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
    ContainerRequest: req,
    Started: true,
    Logger: log.Default(),
  })
  if err != nil {
    return nil, err
  }

  ip, err := container.Host(ctx)
  if err != nil {
    return nil, err
  }

  mappedPort, err := container.MappedPort(ctx, nat.Port(parsedPort))
   if err != nil {
    return nil, err
  }

  uri := fmt.Sprintf("http://%s:%s", ip, mappedPort.Port())

  return &ServerContainer{ Container: container, URI: uri }, nil

}

func TestNewChordNetwork(t *testing.T) {
  ctx := context.Background()
  parentNodePort := 8000
  ports := []int{8001}
  servers := []ServerContainer{}

  _, err := setupServer(ctx, parentNodePort, "")
  if err != nil {
    t.Fatal(err)
  }

  for _, port := range ports {
    parentAddress := fmt.Sprintf("0.0.0.0:%d", parentNodePort)
    srv, err := setupServer(ctx, port, parentAddress)
    if err != nil {
      t.Fatal(err)
    }
    servers = append(servers, *srv)
  }

  defer func() {
    for _, srv := range servers {
      srv.Terminate(ctx)
    }
  }()

  time.Sleep(30 * time.Second)
}
