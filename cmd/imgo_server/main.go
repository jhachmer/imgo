package main

import (
	"github.com/jhachmer/imgo/internal/api"
	"github.com/jhachmer/imgo/internal/config"
)

func main() {
	apiSvr := api.NewServer(config.Envs.Port, config.V1PREFIX, nil)
	apiSvr.Serve()
}
