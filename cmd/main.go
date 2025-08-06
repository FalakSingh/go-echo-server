package main

import "echo-server/cmd/server"

func main() {
	srvr := server.NewServer()
	srvr.Start()
}
