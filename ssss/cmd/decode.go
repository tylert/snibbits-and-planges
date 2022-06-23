package cmd

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/simonfrey/s4"
	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		decodeShares(args)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}

func decodeShares(args []string) {
	sbs := make([][]byte, len(args))

	for i, v := range args {
		data, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			log.Fatal("ERROR: decode - ", err)
		}
		sbs[i] = data
	}

	result, err := s4.RecoverBytes(sbs)
	if err == nil {
		fmt.Println("---")
		fmt.Println(string(result))
		fmt.Println("---")
	} else {
		log.Fatal("ERROR: decode - ", err)
	}
}
