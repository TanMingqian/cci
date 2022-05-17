package docker

import (
	cli "github.com/docker/cli/cli/command"
)

type DockerCli struct {
	cli.Cli
}

func NewDockerCli(ops ...cli.DockerCliOption) (*DockerCli, error) {
	dockerCli, err := cli.NewDockerCli(ops...)
	if err != nil {
		return nil, err
	}
	return &DockerCli{
		dockerCli,
	}, nil
}
