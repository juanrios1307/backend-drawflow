package main

import (
	"backend/routes"
	"backend/server"
)

func main() {

	mux := routes.Route()
	server := server.NewServer(mux)
	server.Run()
}