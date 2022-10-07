package helpertest

import (
	"context"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
)

// ContainerRequest simplifies the params given to CreateTestContainer, supplying its input requirement to generate
// the desired container and an easy way to request which port was effectively mapped.
type ContainerRequest struct {
	Request   testcontainers.ContainerRequest
	PortToMap string
}

// TestContainer combines the requested container return and the effectively mapped port.
type TestContainer struct {
	Container  testcontainers.Container
	MappedPort nat.Port
}

// CreateTestContainer is used to create docker containers that are to be used only for functional tests. Always remember to set its related Envs and terminate with a defer
func CreateTestContainer(ctx context.Context, t *testing.T, containerRequest ContainerRequest) (testContainer TestContainer) {
	container, err := startContainer(ctx, containerRequest.Request)
	if err != nil {
		t.Fatalf("[ helper ] failed to start container: %v\n", err)
	}

	port, err := container.MappedPort(ctx, nat.Port(containerRequest.PortToMap))
	if err != nil {
		t.Fatalf("[ helper ] failed to get mapped port: %v\n", err)
	}

	testContainer = TestContainer{
		Container:  container,
		MappedPort: port,
	}

	return
}

func startContainer(ctx context.Context, request testcontainers.ContainerRequest) (testcontainers.Container, error) {
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
}
