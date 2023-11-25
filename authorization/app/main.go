package main

import (
	src "authorization/core/infrastructure"
)

func main() {
	server := &src.Server{}
	server.Start()
}
