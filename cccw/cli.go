package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/ian-kent/envconf"
	"github.com/vharitonsky/iniflags"
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
	iniflags.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}
