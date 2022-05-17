package client

import (
	"github.com/tanmingqian/cci/client/docker"
)

type Client struct {
	*docker.DockerCli
}
