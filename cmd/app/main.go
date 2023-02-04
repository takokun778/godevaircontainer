package main

import (
	"godevaircontainer/internal/server"
	"os"
)

func main() {
	server.New(os.Getenv("PORT")).Run()
}
