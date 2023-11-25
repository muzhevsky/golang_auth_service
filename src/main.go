package main

import "golang-app/src/server"

func main() {
	server := &server.Server{}
	server.Start()
}
