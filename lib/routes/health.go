package routes

import (
	"fmt"
	"io"
	"net/http"
)

func Check(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /health request\n")
	io.WriteString(w, "OK")
}
