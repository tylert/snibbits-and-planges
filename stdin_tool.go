package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	fmt.Println("Moooo!")
	fmt.Println(message)
}

// https://sj14.gitlab.io/post/2019/02-10-go-stdin-pipe-arg/
