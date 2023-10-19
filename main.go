package main

import (
	"github.com/kaikeventura/cat-runner/src/configuration/server"
)

func main() {
	startServer()
}

func startServer() {
	serverGin := server.NewServer()
	serverGin.Run()
}
