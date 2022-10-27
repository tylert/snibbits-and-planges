package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/ian-kent/envconf"
	"github.com/sethvargo/go-diceware/diceware"
	"github.com/tyler-smith/go-bip39"
	"github.com/vharitonsky/iniflags"
)

// Command-line arguments
var (
	aVersion bool
)

func init() {
	// Help for command-line arguments
	const (
		sVersion = "Display build version information (default false)"
	)

	flag.BoolVar(&aVersion, "version", FromEnvP("CCCW_VERSION", false).(bool), sVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("CCCW_VERSION", false).(bool), sVersion)
	iniflags.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	// Print out the version information
	if aVersion {
		fmt.Println(GetVersion())
		os.Exit(0)
	}

	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Fatal(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatal(err)
	}
	secret_list, err := diceware.Generate(12)
	if err != nil {
		log.Fatal(err)
	}
	secret := strings.Join(secret_list, " ")
	// seed := bip39.NewSeed(mnemonic, secret)

	fmt.Println(mnemonic)
	fmt.Println(secret)
}
