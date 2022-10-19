package main

import (
	"errors"
	"flag"
	"fmt"
	// "log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/google/uuid"
	. "github.com/ian-kent/envconf"
	// "github.com/tv42/base58"
	"github.com/vharitonsky/iniflags"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	// uuid2 "github.com/uuid6/uuid6go-proto"
)

// Command-line arguments
var (
	aDomain    string
	aId        string
	aLong      bool
	aName      string
	aNamespace string
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
		sDomain    = "Domain to use for the UUIDv2 value (Person, Group, Org)"
		sId        = "ID to use for the UUIDv2 value"
		sLong      = "Show the long UUID instead of the short one (default false)"
		sName      = "Name to use for the UUIDv5 or v3 hash"
		sNamespace = "Namespace to use for the UUIDv5 or v3 hash (DNS, OID, URL, X500)"
		sQrFile    = "Also output the UUID as a QR code to a specified JPEG file"
		sQrTerm    = "Also output the UUID as a QR code to the terminal"
		sTypeUuid  = "Generate a new UUID of version (type) v5/v4/v3/v2/v1"
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

// go build -ldflags "-X main.Version=$(git describe --always --dirty --tags)"
var Version string

func GetVersion() string {
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
		fmt.Println(GetVersion())
		os.Exit(0)
	}

	enc := base58Encoder{}
	var (
		err error
		luu string
		suu string
		u64 uint64
		u32 uint32
	)

	// Lengthen or shorten an existing UUID or generate a new one
	if aUuid != "" {
		auu, err := uuid.Parse(aUuid)

		// It might be a short UUID already
		if err != nil {
			auu, err = enc.Decode(aUuid)
		}
		luu = auu.String()
	} else {
		// A non-empty name but default type means we probably want UUIDv5
		if aName != "" && aTypeUuid == "4" {
			aTypeUuid = "5"
		}

		// Get the base10 uint32 value from the id string (always runs since default id is "0")
		if aId != "" {
			u64, err = strconv.ParseUint(aId, 10, 32)
			u32 = uint32(u64)

			if err != nil {
				fmt.Println("Error parsing uint32 string")
				os.Exit(2)
			}
		}

		switch aTypeUuid {
		case "1":
			luu, err = Genv1()
		case "2":
			luu, err = Genv2(aDomain, u32)
		case "3":
			luu, err = Genv3(aName, aNamespace)
		case "4":
			luu, err = Genv4()
		case "5":
			luu, err = Genv5(aName, aNamespace)
		// case "6":
		// 	luu, err = Genv6()
		// case "7":
		// 	luu, err = Genv7()
		// case "8":
		// 	luu, err = Genv8()
		default:
			fmt.Println("Unsupported UUID version")
			os.Exit(3)
		}
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	auu, _ := uuid.Parse(luu)
	suu = enc.Encode(auu)
	if aLong {
		fmt.Println(luu)
	} else {
		fmt.Println(suu)
	}
	if aXtra {
		auu, _ := uuid.Parse(luu)
		fmt.Println(xtraInfo(auu))
	}

	if aQrTerm {
		fmt.Println("go-qrcode terminal stuff doesn't work!!!")
	}

	if aQrFile != "" {
		qrc, err := qrcode.New(suu)
		if err != nil {
			panic(err)
		}
		w, err := standard.New(aQrFile)
		if err != nil {
			panic(err)
		}
		if err = qrc.Save(w); err != nil {
			panic(err)
		}
	}
}

// Other UUID versions???
//   https://github.com/uuid6/uuid6go-proto
//   https://pkg.go.dev/github.com/google/uuid
//   https://github.com/google/uuid
//   https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format-04
//   https://datatracker.ietf.org/doc/html/rfc4122
//   https://en.wikipedia.org/wiki/Universally_unique_identifier

// Other alphabets???
//   https://datatracker.ietf.org/doc/html/draft-msporny-base58-03
//   https://stackoverflow.com/questions/41996761/golang-number-base-conversion/48362821#48362821
//   https://github.com/tv42/base58
//   https://github.com/tv42/zbase32
//   https://github.com/Dasio/base45/blob/main/base45.go
//   https://cs.opensource.google/go/go/+/master:src/encoding/base64/base64.go
//   https://www.rfc-editor.org/info/rfc9285
// base57    '23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
// base58    '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
// base62    '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
// base64    '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/'
// base64url '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_'

// Other features???
//   https://paulgorman.org/technical/blog/20171113164018.html  JSON config
//   https://github.com/signintech/gopdf  generating a PDF

func Genv1() (string, error) {
	// uuid.SetClockSequence(-1)
	// uuid.SetNodeID([]byte{00, 00, 00, 00, 00, 01})
	// uuid.SetNodInterface("")
	uu, err := uuid.NewUUID()
	return uu.String(), err
}

func Genv2(domain string, id uint32) (string, error) {
	// uuid.SetNodeID([]byte{00, 00, 00, 00, 00, 01})
	// uuid.SetNodInterface("")
	switch strings.ToLower(domain) {
	case "person":
		uu, err := uuid.NewDCESecurity(uuid.Person, id)
		return uu.String(), err
	case "group":
		uu, err := uuid.NewDCESecurity(uuid.Group, id)
		return uu.String(), err
	case "org":
		uu, err := uuid.NewDCESecurity(uuid.Org, id)
		return uu.String(), err
	default:
		return uuid.Nil.String(), errors.New("Unsupported domain")
	}
}

func Genv3(name string, namespace string) (string, error) {
	switch strings.ToUpper(namespace) {
	case "DNS":
		uu := uuid.NewMD5(uuid.NameSpaceDNS, []byte(name))
		return uu.String(), nil
	case "OID":
		uu := uuid.NewMD5(uuid.NameSpaceOID, []byte(name))
		return uu.String(), nil
	case "URL":
		uu := uuid.NewMD5(uuid.NameSpaceURL, []byte(name))
		return uu.String(), nil
	case "X500":
		uu := uuid.NewMD5(uuid.NameSpaceX500, []byte(name))
		return uu.String(), nil
	default:
		return uuid.Nil.String(), errors.New("Unsupported namespace")
	}
}

func Genv4() (string, error) {
	uu, err := uuid.NewRandom()
	return uu.String(), err
}

func Genv5(name string, namespace string) (string, error) {
	switch strings.ToUpper(namespace) {
	case "DNS":
		uu := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(name))
		return uu.String(), nil
	case "OID":
		uu := uuid.NewSHA1(uuid.NameSpaceOID, []byte(name))
		return uu.String(), nil
	case "URL":
		uu := uuid.NewSHA1(uuid.NameSpaceURL, []byte(name))
		return uu.String(), nil
	case "X500":
		uu := uuid.NewSHA1(uuid.NameSpaceX500, []byte(name))
		return uu.String(), nil
	default:
		return uuid.Nil.String(), errors.New("Unsupported namespace")
	}
}

func xtraInfo(luu uuid.UUID) string {
	ver := strings.Split(luu.Version().String(), "_")
	output := fmt.Sprintf("UUID Version:%s Variant:%s", ver[1], luu.Variant())

	switch ver[1] {
	case "2":
		output += fmt.Sprintf(" Domain:%s Id:%s", luu.Domain().String(), luu.ID())
	}

	return output
}

type base58Encoder struct{}

func (enc base58Encoder) Encode(luu uuid.UUID) string {
	return base58.Encode(luu[:])
}

func (enc base58Encoder) Decode(suu string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(suu))
}
