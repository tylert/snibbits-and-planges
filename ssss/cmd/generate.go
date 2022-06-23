package cmd

import (
	"encoding/base64"
	"fmt"
	"log"
	"strconv"

	"github.com/simonfrey/s4"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateShares(args)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateShares(args []string) {
	if len(args) != 3 {
		log.Fatal("ERROR: generate - must have 3 args")
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("ERROR: generate - Expecting integer: ", num)
	}
	if num < 3 {
		log.Fatal("ERROR: generate - At least 3 shares required, found ", num)
	}

	req, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("ERROR: generate - Expecting integer: ", req)
	}

	fmt.Printf("shares: %d;  required: %d\n", num, req)

	message := []byte(args[2])
	shares, err := s4.DistributeBytes(message, uint64(num), uint64(req))
	if err != nil {
		log.Fatal("ERROR: generate - ERROR: ", err)
	}

	fmt.Println("---")
	for _, share := range shares {
		fmt.Println(base64.StdEncoding.EncodeToString([]byte(share)))
	}
	fmt.Println("---")
}
