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
	aId        string
	aLong      bool
	aName      string
	aNamespace string
	aNodeId    string
	aQrFile    string
	aQrTerm    bool
	aTypeUuid  string
	aUuid      string
	aVersion   bool
	aXtra      bool
)

func init() {
	// Help for command-line arguments
	const (
		sDomain    = "Domain to use for UUIDv2/v1 - Person/Group/Org"
		sId        = "ID to use for UUIDv2/v1"
		sLong      = "Show the long UUID instead of the short one (default false)"
		sName      = "Name to use for the UUIDv5/v3 hash"
		sNamespace = "Namespace to use for UUIDv5/v3 hash - DNS/OID/URL/X500"
		sNodeId    = "NodeID [interface] to use for UUIDv2/v1 - random if 'none'"
		sQrFile    = "Also output the UUID as a QR code to a specified JPEG file"
		sQrTerm    = "Also output the UUID as a QR code to the terminal"
		sTypeUuid  = "Version [type] of UUID to generate - UUIDv5/v4/v3/v2/v1"
		sUuid      = "Existing UUID to shorten or lengthen"
		sVersion   = "Display build version information (default false)"
		sXtra      = "Show extra details about the UUID (default false)"
	)

	flag.StringVar(&aDomain, "domain", FromEnvP("SHORTUUID_DOMAIN", "Person").(string), sDomain)
	flag.StringVar(&aDomain, "d", FromEnvP("SHORTUUID_DOMAIN", "Person").(string), sDomain)
	flag.StringVar(&aId, "id", FromEnvP("SHORTUUID_ID", "0").(string), sId)
	flag.StringVar(&aId, "i", FromEnvP("SHORTUUID_ID", "0").(string), sId)
	flag.BoolVar(&aLong, "long", FromEnvP("SHORTUUID_LONG", false).(bool), sLong)
	flag.BoolVar(&aLong, "l", FromEnvP("SHORTUUID_LONG", false).(bool), sLong)
	flag.StringVar(&aName, "name", FromEnvP("SHORTUUID_NAME", "").(string), sName)
	flag.StringVar(&aName, "n", FromEnvP("SHORTUUID_NAME", "").(string), sName)
	flag.StringVar(&aNamespace, "namespace", FromEnvP("SHORTUUID_NAMESPACE", "DNS").(string), sNamespace)
	flag.StringVar(&aNamespace, "ns", FromEnvP("SHORTUUID_NAMESPACE", "DNS").(string), sNamespace)
	flag.StringVar(&aNodeId, "nodeid", FromEnvP("SHORTUUID_NODEID", "").(string), sNodeId)
	flag.StringVar(&aNodeId, "o", FromEnvP("SHORTUUID_NODEID", "").(string), sNodeId)
	flag.StringVar(&aQrFile, "qrfile", FromEnvP("SHORTUUID_QRFILE", "").(string), sQrFile)
	flag.StringVar(&aQrFile, "qf", FromEnvP("SHORTUUID_QRFILE", "").(string), sQrFile)
	flag.BoolVar(&aQrTerm, "qrterm", FromEnvP("SHORTUUID_QRTERM", false).(bool), sQrTerm)
	flag.BoolVar(&aQrTerm, "qt", FromEnvP("SHORTUUID_QRTERM", false).(bool), sQrTerm)
	flag.StringVar(&aTypeUuid, "typeuuid", FromEnvP("SHORTUUID_TYPEUUID", "4").(string), sTypeUuid)
	flag.StringVar(&aTypeUuid, "t", FromEnvP("SHORTUUID_TYPEUUID", "4").(string), sTypeUuid)
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
