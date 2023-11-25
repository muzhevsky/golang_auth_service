package main

import "golang-app/src/myServer"

func main() {
	server := &myServer.Server{}
	server.Start()
}
