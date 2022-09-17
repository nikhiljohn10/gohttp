package main

import (
	"fmt"
	"gohttp/server"
)

func main() {
	fmt.Println("Running http server at http://localhost:8080")
	server.Run()
}
