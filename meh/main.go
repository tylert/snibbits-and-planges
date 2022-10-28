package main

import (
	// "encoding/csv"
	// "encoding/json"
	"fmt"
	"os"
	// "github.com/goccy/go-yaml"
	// "github.com/jszwec/csvutil"
)

func main() {
	// Print out the version information
	if aVersion {
		fmt.Println(GetVersion())
		os.Exit(0)
	}

	// if aImport != "" {
	//   file, err := os.Open(aImport)
	//   if err != nil {
	//     panic(err)
	//   }

	//   reader := csv.NewReader(file)
	//   records, _ := reader.ReadAll()
	//   fmt.Println(records)
	// }

	// bs, err := json.MarshalIndent(foo, "", "  ")
	// if err != nil {
	//   panic(err)
	// }
	// fmt.Println(string(bs))
}
