package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", moo)

	fmt.Println(":8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func moo(w http.ResponseWriter, r *http.Request) {
	requesterIpAndPort := r.RemoteAddr
	requesterIp := r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")]

	fmt.Println(requesterIpAndPort)
	w.Write([]byte(requesterIp + "\n"))
}
