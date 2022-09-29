package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	// "github.com/goccy/go-yaml"
	. "github.com/ian-kent/envconf"
	// "github.com/vharitonsky/iniflags"
)

// Command-line arguments
var (
	aExport  string
	aImport  string
	aVersion bool
)

func init() {
	// Help for command-line arguments
	const (
		sExport  = "Export YAML or JSON to a file"
		sImport  = "Import YAML or JSON from a file"
		sVersion = "Display build version information (default false)"
	)

	flag.StringVar(&aExport, "export", FromEnvP("MOOYAML_EXPORT", "").(string), sExport)
	flag.StringVar(&aExport, "e", FromEnvP("MOOYAML_EXPORT", "").(string), sExport)
	flag.StringVar(&aImport, "import", FromEnvP("MOOYAML_IMPORT", "").(string), sImport)
	flag.StringVar(&aImport, "i", FromEnvP("MOOYAML_IMPORT", "").(string), sImport)
	flag.BoolVar(&aVersion, "version", FromEnvP("MOOYAML_VERSION", false).(bool), sVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("MOOYAML_VERSION", false).(bool), sVersion)
	// iniflags.Parse()
	flag.Parse()

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
}
