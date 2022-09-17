package main

import (
	"errors"
	"example/routes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", routes.getRoot)
	http.HandleFunc("/hello", routes.getHello)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
