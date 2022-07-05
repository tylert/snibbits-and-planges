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

// XXX FIXME TODO  Allow it to work on something other than 127.0.0.1 or ::1
// https://gist.github.com/2minchul/191716b3ca8799f53362746731d08e91
// https://gist.github.com/bacher09/51ce161105a9e1f49b8b917f8eccd3c5
