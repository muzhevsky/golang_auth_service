package main

import (
	src "authorization/core"
)

func main() {
	server := &src.Server{}
	server.Start()
}
