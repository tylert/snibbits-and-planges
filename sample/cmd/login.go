package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		verbosity, _ = cmd.Flags().GetCount("verbose")
		doLogin(args)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().CountP("verbose", "v", "counted verbosity")
}

func doLogin(args []string) {
	fmt.Println("Login called")
	if verbosity > 0 {
		fmt.Println("args: " + strings.Join(args, " "))
	}
}
