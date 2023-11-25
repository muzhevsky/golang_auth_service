package main

import (
	src "authorization/internal"
)

func main() {
	server := &src.Server{}
	server.Start()
}
