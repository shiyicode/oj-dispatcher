package docker

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

func newClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}
	return cli, err
}

func CreateContainer(image string, cmd []string) (string, error) {
	cli, err := newClient()
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	containerBody, err := cli.ContainerCreate(ctx,
		&container.Config{
			Image: image,
			Cmd:   cmd,
			User:  "root",
		}, &container.HostConfig{
			Resources: container.Resources{
				NanoCPUs: 2,
				Memory:   512000000,
			},
		}, nil, "")
	if err != nil {
		log.WithField("err", err.Error()).Error("docker container create failure")
		return "", err
	}
	return containerBody.ID, nil
}

func StartContainer(containerId string) error {
	cli, err := newClient()
	if err != nil {
		return err
	}
	ctx := context.Background()

	err = cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{})
	return err
}

func StopContainer(containerID string) error {
	cli, err := newClient()
	if err != nil {
		return err
	}
	timeout := time.Second * 10
	err = cli.ContainerStop(context.Background(), containerID, &timeout)
	return err
}

func RemoveContainer(containerID string, force bool, removeVolumes bool, removeLinks bool) error {
	cli, err := newClient()
	if err != nil {
		return err
	}
	ctx := context.Background()

	options := types.ContainerRemoveOptions{Force: force, RemoveVolumes: removeVolumes, RemoveLinks: removeLinks}
	if err := cli.ContainerRemove(ctx, containerID, options); err != nil {
		return err
	}
	return nil
}

func KillContainer(containerId string) error {
	cli, err := newClient()
	if err != nil {
		return err
	}
	ctx := context.Background()

	err = cli.ContainerKill(ctx, containerId, "SIGKILL")
	return err
}

func ListContainers() ([]types.Container, error) {
	cli, err := newClient()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func ListImages() ([]types.ImageSummary, error) {
	cli, err := newClient()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return nil, nil
	}

	return images, nil
}

func PullImage(imageName string) error {
	cli, err := newClient()
	if err != nil {
		return err
	}
	ctx := context.Background()

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(os.Stdout, out)

	return nil
}

func PrintLogContainer(containerID string) error {
	cli, err := newClient()
	if err != nil {
		return err
	}
	ctx := context.Background()

	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, containerID, options)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	io.Copy(os.Stdout, out)

	return nil
}
