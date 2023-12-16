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
	// Usage for command-line arguments
	const (
		// uClockSeq  = "Clock sequence [14-bit number] to use for UUIDv6/v2/v1"
		uDomain    = "Domain [PERSON/GROUP/ORG] to use for UUIDv2"
		uEncoding  = "Encoding [BASE58/NONE] to use for shortening UUID"
		uId        = "ID to use for UUIDv2"
		uName      = "Name to use for UUIDv5/v3"
		uNamespace = "Namespace [DNS/OID/URL/X500] to use for UUIDv5/v3"
		uNode      = "Node [48-bit MAC] to use for UUIDv6/v2/v1"
		uQrFile    = "Also output the UUID as a QR code to a specified JPEG file"
		uQrTerm    = "Also output the UUID as a QR code to the terminal"
		uType      = "Type (version) [UUIDv6/v5/v4/v3/v2/v1] of UUID to generate"
		uUuid      = "Existing UUID to shorten or lengthen"
		uVersion   = "Display build version information (default false)"
		uXtra      = "Show extra details about the UUID (default false)"
	)

	// flag.StringVar(&aClockSeq, "clockseq", FromEnvP("SHORTUUID_CLOCKSEQ", "").(string), uClockSeq)
	// flag.StringVar(&aClockSeq, "c", FromEnvP("SHORTUUID_CLOCKSEQ", "").(string), uClockSeq)
	flag.StringVar(&aDomain, "domain", FromEnvP("SHORTUUID_DOMAIN", "PERSON").(string), uDomain)
	flag.StringVar(&aDomain, "d", FromEnvP("SHORTUUID_DOMAIN", "PERSON").(string), uDomain)
	flag.StringVar(&aEncoding, "encoding", FromEnvP("SHORTUUID_ENCODING", "BASE58").(string), uEncoding)
	flag.StringVar(&aEncoding, "e", FromEnvP("SHORTUUID_ENCODING", "BASE58").(string), uEncoding)
	flag.StringVar(&aId, "id", FromEnvP("SHORTUUID_ID", "0").(string), uId)
	flag.StringVar(&aId, "i", FromEnvP("SHORTUUID_ID", "0").(string), uId)
	flag.StringVar(&aName, "name", FromEnvP("SHORTUUID_NAME", "").(string), uName)
	flag.StringVar(&aName, "n", FromEnvP("SHORTUUID_NAME", "").(string), uName)
	flag.StringVar(&aNamespace, "namespace", FromEnvP("SHORTUUID_NAMESPACE", "DNS").(string), uNamespace)
	flag.StringVar(&aNamespace, "ns", FromEnvP("SHORTUUID_NAMESPACE", "DNS").(string), uNamespace)
	flag.StringVar(&aNode, "node", FromEnvP("SHORTUUID_NODE", "").(string), uNode)
	flag.StringVar(&aNode, "o", FromEnvP("SHORTUUID_NODE", "").(string), uNode)
	flag.StringVar(&aQrFile, "qrfile", FromEnvP("SHORTUUID_QRFILE", "").(string), uQrFile)
	flag.StringVar(&aQrFile, "qf", FromEnvP("SHORTUUID_QRFILE", "").(string), uQrFile)
	flag.BoolVar(&aQrTerm, "qrterm", FromEnvP("SHORTUUID_QRTERM", false).(bool), uQrTerm)
	flag.BoolVar(&aQrTerm, "qt", FromEnvP("SHORTUUID_QRTERM", false).(bool), uQrTerm)
	flag.StringVar(&aType, "type", FromEnvP("SHORTUUID_TYPE", "4").(string), uType)
	flag.StringVar(&aType, "t", FromEnvP("SHORTUUID_TYPE", "4").(string), uType)
	flag.StringVar(&aUuid, "uuid", FromEnvP("SHORTUUID_UUID", "").(string), uUuid)
	flag.StringVar(&aUuid, "u", FromEnvP("SHORTUUID_UUID", "").(string), uUuid)
	flag.BoolVar(&aVersion, "version", FromEnvP("SHORTUUID_VERSION", false).(bool), uVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("SHORTUUID_VERSION", false).(bool), uVersion)
	flag.BoolVar(&aXtra, "xtra", FromEnvP("SHORTUUID_XTRA", false).(bool), uXtra)
	flag.BoolVar(&aXtra, "x", FromEnvP("SHORTUUID_XTRA", false).(bool), uXtra)

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
	// cfg, err := ini.LooseLoad(fmt.Sprintf("%s/.config/shortuuid/defaults", home))
	// https://ini.unknwon.io/docs

	flag.Parse()
	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}
