package routes

import (
	"fmt"
	"gohttp/lib/hi"
	"io"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	msg := hi.Hello("World!")
	fmt.Printf("got /hello request\n")
	io.WriteString(w, msg+"\n")
}
