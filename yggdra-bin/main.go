package main

import (
	"yggdra/server"
)

func main() {
	port := "9898"
	listenAdress := ":" + port

	server.Serve(listenAdress)
}
