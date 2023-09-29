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
	//aClockSeq  string
	aDomain    string
	aEncoding  string
	aId        string
	aName      string
	aNamespace string
	aNode      string
	aQrFile    string
	aQrTerm    bool
	aType      string
	aUuid      string
	aVersion   bool
	aXtra      bool
)

func init() {
	// Help for command-line arguments
	const (
		// sClockSeq  = "Clock sequence [14-bit number] to use for UUIDv2/v1"
		sDomain    = "Domain [PERSON/GROUP/ORG] to use for UUIDv2"
		sEncoding  = "Encoding [BASE58/NONE] to use for shortening UUID"
		sId        = "ID to use for UUIDv2"
		sName      = "Name to use for UUIDv5/v3"
		sNamespace = "Namespace [DNS/OID/URL/X500] to use for UUIDv5/v3"
		sNode      = "Node [48-bit MAC] to use for UUIDv2/v1"
		sQrFile    = "Also output the UUID as a QR code to a specified JPEG file"
		sQrTerm    = "Also output the UUID as a QR code to the terminal"
		sType      = "Type (version) [UUIDv5/v4/v3/v2/v1] of UUID to generate"
		sUuid      = "Existing UUID to shorten or lengthen"
		sVersion   = "Display build version information (default false)"
		sXtra      = "Show extra details about the UUID (default false)"
	)

	// flag.StringVar(&aClockSeq, "clockseq", FromEnvP("SHORTUUID_CLOCKSEQ", "").(string), sClockSeq)
	// flag.StringVar(&aClockSeq, "c", FromEnvP("SHORTUUID_CLOCKSEQ", "").(string), sClockSeq)
	flag.StringVar(&aDomain, "domain", FromEnvP("SHORTUUID_DOMAIN", "PERSON").(string), sDomain)
	flag.StringVar(&aDomain, "d", FromEnvP("SHORTUUID_DOMAIN", "PERSON").(string), sDomain)
	flag.StringVar(&aEncoding, "encoding", FromEnvP("SHORTUUID_ENCODING", "BASE58").(string), sEncoding)
	flag.StringVar(&aEncoding, "e", FromEnvP("SHORTUUID_ENCODING", "BASE58").(string), sEncoding)
	flag.StringVar(&aId, "id", FromEnvP("SHORTUUID_ID", "0").(string), sId)
	flag.StringVar(&aId, "i", FromEnvP("SHORTUUID_ID", "0").(string), sId)
	flag.StringVar(&aName, "name", FromEnvP("SHORTUUID_NAME", "").(string), sName)
	flag.StringVar(&aName, "n", FromEnvP("SHORTUUID_NAME", "").(string), sName)
	flag.StringVar(&aNamespace, "namespace", FromEnvP("SHORTUUID_NAMESPACE", "DNS").(string), sNamespace)
	flag.StringVar(&aNamespace, "ns", FromEnvP("SHORTUUID_NAMESPACE", "DNS").(string), sNamespace)
	flag.StringVar(&aNode, "node", FromEnvP("SHORTUUID_NODE", "").(string), sNode)
	flag.StringVar(&aNode, "o", FromEnvP("SHORTUUID_NODE", "").(string), sNode)
	flag.StringVar(&aQrFile, "qrfile", FromEnvP("SHORTUUID_QRFILE", "").(string), sQrFile)
	flag.StringVar(&aQrFile, "qf", FromEnvP("SHORTUUID_QRFILE", "").(string), sQrFile)
	flag.BoolVar(&aQrTerm, "qrterm", FromEnvP("SHORTUUID_QRTERM", false).(bool), sQrTerm)
	flag.BoolVar(&aQrTerm, "qt", FromEnvP("SHORTUUID_QRTERM", false).(bool), sQrTerm)
	flag.StringVar(&aType, "type", FromEnvP("SHORTUUID_TYPE", "4").(string), sType)
	flag.StringVar(&aType, "t", FromEnvP("SHORTUUID_TYPE", "4").(string), sType)
	flag.StringVar(&aUuid, "uuid", FromEnvP("SHORTUUID_UUID", "").(string), sUuid)
	flag.StringVar(&aUuid, "u", FromEnvP("SHORTUUID_UUID", "").(string), sUuid)
	flag.BoolVar(&aVersion, "version", FromEnvP("SHORTUUID_VERSION", false).(bool), sVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("SHORTUUID_VERSION", false).(bool), sVersion)
	flag.BoolVar(&aXtra, "xtra", FromEnvP("SHORTUUID_XTRA", false).(bool), sXtra)
	flag.BoolVar(&aXtra, "x", FromEnvP("SHORTUUID_XTRA", false).(bool), sXtra)

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
	iniflags.SetAllowMissingConfigFile(true)
	home, _ := os.UserHomeDir()
	iniflags.SetConfigFile(fmt.Sprintf("%s/.config/shortuuid/defaults", home))
	iniflags.Parse() // Replace with flag.Parse() eventually?!?
	// "gopkg.in/ini.v1"
	// https://github.com/go-ini/ini

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}
