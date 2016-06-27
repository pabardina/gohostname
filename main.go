package main

import (
	"io"
	"net/http"
	"os"
)

func getHostame(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		return
	}
	io.WriteString(w, "Hostname: "+name)
}

func main() {
	http.HandleFunc("/", getHostame)
	http.ListenAndServe(":8000", nil)
}
