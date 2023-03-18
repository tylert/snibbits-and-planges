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

	flag.StringVar(&aExport, "export", FromEnvP("MEH_EXPORT", "").(string), sExport)
	flag.StringVar(&aExport, "e", FromEnvP("MEH_EXPORT", "").(string), sExport)
	flag.StringVar(&aImport, "import", FromEnvP("MEH_IMPORT", "").(string), sImport)
	flag.StringVar(&aImport, "i", FromEnvP("MEH_IMPORT", "").(string), sImport)
	flag.BoolVar(&aVersion, "version", FromEnvP("MEH_VERSION", false).(bool), sVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("MEH_VERSION", false).(bool), sVersion)
	iniflags.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}
