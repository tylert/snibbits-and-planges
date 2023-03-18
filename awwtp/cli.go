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
	aCount   int
	aDigits  int
	aSeed    string
	aType    string
	aVersion bool
)

func init() {
	// Help for command-line arguments
	const (
		sCount   = "Number of seconds to countdown between tokens"
		sDigits  = "Number of digits to output for the tokens"
		sSeed    = "Seed value to use as input"
		sType    = "Hash algorithm to use SHA1/SHA256/SHA512 (default SHA1)"
		sVersion = "Display build version information (default false)"
	)

	flag.IntVar(&aCount, "countdown", FromEnvP("AWWTP_COUNT", 30).(int), sCount)
	flag.IntVar(&aCount, "c", FromEnvP("AWWTP_COUNT", 30).(int), sCount)
	flag.IntVar(&aDigits, "digits", FromEnvP("AWWTP_DIGITS", 6).(int), sDigits)
	flag.IntVar(&aDigits, "d", FromEnvP("AWWTP_DIGITS", 6).(int), sDigits)
	flag.StringVar(&aSeed, "seed", FromEnvP("AWWTP_SEED", "").(string), sSeed)
	flag.StringVar(&aSeed, "s", FromEnvP("AWWTP_SEED", "").(string), sSeed)
	flag.StringVar(&aType, "type", FromEnvP("AWWTP_TYPE", "SHA1").(string), sType)
	flag.StringVar(&aType, "t", FromEnvP("AWWTP_TYPE", "SHA1").(string), sType)
	flag.BoolVar(&aVersion, "version", FromEnvP("AWWTP_VERSION", false).(bool), sVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("AWWTP_VERSION", false).(bool), sVersion)
	iniflags.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}
