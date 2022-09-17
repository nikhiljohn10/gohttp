package routes

import (
	"example/hi"
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	msg := hi.Hello("World!")
	fmt.Printf("got /hello request\n")
	io.WriteString(w, msg+"\n")
}
