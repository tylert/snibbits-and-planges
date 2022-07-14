package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	defaultConfigFilename = "sample"
	envPrefix             = "SAMPLE"
)

var version = "0.0.1"
var verbosity = 0

var rootCmd = &cobra.Command{
	Use:     "sample",
	Version: version,
	Short:   "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
to quickly create a Cobra application.`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initializeConfig(cmd)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	v.SetConfigName(defaultConfigFilename)

	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/.config/sample") // TODO: Use XDG??

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()
	bindFlags(cmd, v)

	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			err := v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
			if err != nil {
				return
			}
		}

		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if err != nil {
				return
			}
		}
	})
}
