package server

import (
	"errors"
	"fmt"
	"gohttp/lib/routes"
	"net/http"
	"os"
)

func Run() {
	http.HandleFunc("/", routes.GetRoot)
	http.HandleFunc("/hello", routes.GetHello)
	http.HandleFunc("/health", routes.Check)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
