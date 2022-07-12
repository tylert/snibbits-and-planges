package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("boo no pipe")
	} else {
		var reader = bufio.NewReader(os.Stdin)
		msg, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println(msg)
	}
}

// https://sj14.gitlab.io/post/2019/02-10-go-stdin-pipe-arg/
// https://stackoverflow.com/questions/22744443/check-if-there-is-something-to-read-on-stdin-in-golang
// https://stackoverflow.com/questions/22563616/determine-if-stdin-has-data-with-go
// https://coderwall.com/p/zyxyeg/golang-having-fun-with-os-stdin-and-shell-pipes
