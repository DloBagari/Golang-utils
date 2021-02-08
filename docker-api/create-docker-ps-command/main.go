package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

func main() {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		logrus.Fatal(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		logrus.Fatal(err)
	}
	for _, container := range containers {
		logrus.Infof("Image: %s with name %s %s", container.Image, container.Names, container.Status)
	}
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		logrus.Fatal(err)
	}
	for _, img := range images {
		logrus.Infof("Image: %d with name %d %s", img.Containers, img.Size, img.ID)
	}

}
