package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/ian-kent/envconf"
	//"gopkg.in/ini.v1"
)

// Command-line arguments
var (
	aVersion bool
)

func init() {
	// Usage for command-line arguments
	const (
		uVersion = "Display build version information (default false)"
	)

	flag.BoolVar(&aVersion, "version", FromEnvP("CCCW_VERSION", false).(bool), uVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("CCCW_VERSION", false).(bool), uVersion)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		// flag.VisitAll(func(f *flag.Flag) {
		//   fmt.Fprintf(os.Stderr, "%v %v %v\n", f.Name, f.Value, f.Usage)
		// })
	}

	// FlagSet for sub-commands???
	// https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go

	// Attempt to gracefully load things from a known config file location
	// cfg := ini.Empty()
	// home, _ := os.UserHomeDir()
	// cfg, err := ini.LooseLoad(fmt.Sprintf("%s/.config/yabba/defaults", home))
	// https://ini.unknwon.io/docs

	flag.Parse()
	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}
