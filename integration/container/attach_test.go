package container // import "github.com/docker/docker/integration/container"

import (
	"context"
	"bytes"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/pkg/stdcopy"
	"gotest.tools/v3/assert"
)

func TestAttachWithTTY(t *testing.T) {
	testAttach(t, true, types.MediaTypeRawStream)
}

func TestAttachWithoutTTy(t *testing.T) {
	testAttach(t, false, types.MediaTypeMultiplexedStream)
}

func testAttach(t *testing.T, tty bool, expected string) {
	defer setupTest(t)()
	client := testEnv.APIClient()

	resp, err := client.ContainerCreate(context.Background(),
		&container.Config{
			Image: "busybox",
			Cmd:   []string{"echo", "hello"},
			Tty:   tty,
		},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		nil,
		"",
	)
	assert.NilError(t, err)
	container := resp.ID
	defer client.ContainerRemove(context.Background(), container, types.ContainerRemoveOptions{
		Force: true,
	})

	attach, err := client.ContainerAttach(context.Background(), container, types.ContainerAttachOptions{
		Stdout: true,
		Stderr: true,
	})
	assert.NilError(t, err)
	mediaType, ok := attach.MediaType()
	assert.Check(t, ok)
	assert.Check(t, mediaType == expected)
}

func TestAttachBeforeContainerStart(t *testing.T) {
	defer setupTest(t)()
	client := testEnv.APIClient()

	resp, err := client.ContainerCreate(context.Background(),
		&container.Config{
			Image: "busybox",
			Cmd:   []string{"echo", "-n", "hello"},
			Tty:   false,
		},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		nil,
		"",
	)
	assert.NilError(t, err)
	containerId := resp.ID

	attach, err := client.ContainerAttach(context.Background(), containerId, types.ContainerAttachOptions{
		Stdout: true,
		Stderr: true,
		Stream: true,
	})
	assert.NilError(t, err)

	defer client.ContainerRemove(context.Background(), containerId, types.ContainerRemoveOptions{
		Force: true,
	})

	err = client.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	assert.NilError(t, err)

	statusCh, errCh := client.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		assert.NilError(t, err)
	case <-statusCh:
	}

	var stdout, stderr bytes.Buffer
	_, err = stdcopy.StdCopy(&stdout, &stderr, attach.Reader)
	assert.NilError(t, err)

	stdoutStr := stdout.String()

	// Passes in docker v20.10.18
	assert.Equal(t, "hello", stdoutStr)
}

func TestAttachAfterContainerStart(t *testing.T) {
	defer setupTest(t)()
	client := testEnv.APIClient()

	resp, err := client.ContainerCreate(context.Background(),
		&container.Config{
			Image: "busybox",
			Cmd:   []string{"echo", "-n", "hello"},
			Tty:   false,
		},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		nil,
		"",
	)
	assert.NilError(t, err)
	containerId := resp.ID

	defer client.ContainerRemove(context.Background(), containerId, types.ContainerRemoveOptions{
		Force: true,
	})

	err = client.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	assert.NilError(t, err)

	// Blocks until container is started in docker v20.10.18
	attach, err := client.ContainerAttach(context.Background(), containerId, types.ContainerAttachOptions{
		Stdout: true,
		Stderr: true,
		Stream: true,
	})
	assert.NilError(t, err)

	statusCh, errCh := client.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		assert.NilError(t, err)
	case <-statusCh:
	}

	var stdout, stderr bytes.Buffer
	_, err = stdcopy.StdCopy(&stdout, &stderr, attach.Reader)
	assert.NilError(t, err)

	stdoutStr := stdout.String()

	assert.Equal(t, "hello", stdoutStr)
}
