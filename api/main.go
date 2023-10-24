package main

import (
	"github.com/kaikeventura/cat-runner/src/configuration/container"
	"github.com/kaikeventura/cat-runner/src/configuration/server"
)

func main() {
	staterContainer()
	startServer()
}

func startServer() {
	serverGin := server.NewServer()
	serverGin.Run()
}

func staterContainer() {
	container.BuildConstainer()
}
