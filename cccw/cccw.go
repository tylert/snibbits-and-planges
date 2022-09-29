package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
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

// go build -ldflags "-X main.Version=$(git describe --always --dirty --tags)"
var Version string

func getVersion() string {
	var barch, bos, bmod, brev, btime, suffix string
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			switch setting.Key {
			case "GOARCH":
				barch = setting.Value
			case "GOOS":
				bos = setting.Value
			case "vcs.modified":
				bmod = setting.Value
			case "vcs.revision":
				brev = setting.Value[0:7]
			case "vcs.time":
				btime = setting.Value
			}
		}
	}
	// If we didn't specify a version string, use the git commit
	if Version == "" {
		Version = brev
	}
	// If the git repo wasn't clean, say so in the version string
	if bmod == "true" {
		suffix = "-dirty"
	}
	return fmt.Sprintf("%s%s %s %s %s", Version, suffix, bos, barch, btime)
}

func main() {
	// Print out the version information
	if aVersion {
		fmt.Println(getVersion())
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
