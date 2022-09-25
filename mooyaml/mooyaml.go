package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	// "encoding/json"
	// "github.com/goccy/go-yaml"
)

var (
	aExport  string
	aImport  string
	aVersion bool
)

func init() {
	const (
		sExport  = "Export YAML or JSON to a file"
		sImport  = "Import YAML or JSON from a file"
		sVersion = "Display build version information (default false)"
	)

	flag.StringVar(&aExport, "export", "", sExport)
	flag.StringVar(&aExport, "e", "", sExport)
	flag.StringVar(&aImport, "import", "", sImport)
	flag.StringVar(&aImport, "i", "", sImport)
	flag.BoolVar(&aVersion, "version", false, sVersion)
	flag.BoolVar(&aVersion, "v", false, sVersion)
	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}

// go build -ldflags "-X main.Version=$(git describe --always --dirty --tags)"
var Version string

func printVersion() {
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
	fmt.Println(fmt.Sprintf("%s%s %s %s %s", Version, suffix, bos, barch, btime))
}

func main() {
	// Print out the version information
	if aVersion {
		printVersion()
		os.Exit(0)
	}
}
