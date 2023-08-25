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
		sDomain    = "Domain to use for UUIDv2 - PERSON/GROUP/ORG"
		sEncoding  = "Encoding to use for shortening UUID - BASE58/NONE"
		sId        = "ID to use for UUIDv2"
		sName      = "Name to use for UUIDv5/v3"
		sNamespace = "Namespace to use for UUIDv5/v3 - DNS/OID/URL/X500"
		sNode      = "Node [interface name] to use for UUIDv2/v1 - RANDOM/eth0/etc."
		sQrFile    = "Also output the UUID as a QR code to a specified JPEG file"
		sQrTerm    = "Also output the UUID as a QR code to the terminal"
		sType      = "Version [type] of UUID to generate - UUIDv5/v4/v3/v2/v1"
		sUuid      = "Existing UUID to shorten or lengthen"
		sVersion   = "Display build version information (default false)"
		sXtra      = "Show extra details about the UUID (default false)"
	)

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
	iniflags.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Error: Unused command line arguments detected.\n")
		flag.Usage()
		os.Exit(1)
	}
}
