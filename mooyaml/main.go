package main

import (
	// "encoding/json"
	"fmt"
	"os"
	// "github.com/goccy/go-yaml"
)

func main() {
	// Print out the version information
	if aVersion {
		fmt.Println(GetVersion())
		os.Exit(0)
	}

	// bs, err := json.MarshalIndent(foo, "", "  ")
	// if err != nil {
	//   panic(err)
	// }
	// fmt.Println(string(bs))
}
