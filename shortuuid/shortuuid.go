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
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	// uuid2 "github.com/uuid6/uuid6go-proto"
	// "github.com/nicksnyder/basen"
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
	aUuid      string
	aUuidType  string
	aVersion   bool
	aXtra      bool
)

func init() {
	// Help for command-line arguments
	const (
		sDomain    = "Domain to use for the UUIDv2 value"
		sId        = "ID to use for the UUIDv2 value"
		sLong      = "Show the long UUID instead of the short one (default false)"
		sName      = "Name to use for the UUIDv5 or v3 hash"
		sNamespace = "Namespace to use for the UUIDv5 or v3 hash"
		sQrFile    = "Also output the UUID as a QR code to a specified JPEG file"
		sQrTerm    = "Also output the UUID as a QR code to the terminal"
		sUuid      = "Existing UUID to shorten or lengthen"
		sUuidType  = "Generate a new UUID of version v5/v4/v3/v2/v1"
		sVersion   = "Display build version information (default false)"
		sXtra      = "Show extra details about the UUID (default false)"
	)

	flag.StringVar(&aDomain, "domain", FromEnvP("DOMAIN", "Person").(string), sDomain)
	flag.StringVar(&aDomain, "d", FromEnvP("DOMAIN", "Person").(string), sDomain)
	flag.StringVar(&aId, "id", FromEnvP("ID", "0").(string), sId)
	flag.StringVar(&aId, "i", FromEnvP("ID", "0").(string), sId)
	flag.BoolVar(&aLong, "long", FromEnvP("LONG", false).(bool), sLong)
	flag.BoolVar(&aLong, "l", FromEnvP("LONG", false).(bool), sLong)
	flag.StringVar(&aName, "name", FromEnvP("NAME", "").(string), sName)
	flag.StringVar(&aName, "n", FromEnvP("NAME", "").(string), sName)
	flag.StringVar(&aNamespace, "namespace", FromEnvP("NAMESPACE", "DNS").(string), sNamespace)
	flag.StringVar(&aNamespace, "ns", FromEnvP("NAMESPACE", "DNS").(string), sNamespace)
	flag.StringVar(&aQrFile, "qrfile", FromEnvP("QRFILE", "").(string), sQrFile)
	flag.StringVar(&aQrFile, "qf", FromEnvP("QRFILE", "").(string), sQrFile)
	flag.BoolVar(&aQrTerm, "qrterm", FromEnvP("QRTERM", false).(bool), sQrTerm)
	flag.BoolVar(&aQrTerm, "qt", FromEnvP("QRTERM", false).(bool), sQrTerm)
	flag.StringVar(&aUuid, "uuid", FromEnvP("UUID", "").(string), sUuid)
	flag.StringVar(&aUuid, "u", FromEnvP("UUID", "").(string), sUuid)
	flag.StringVar(&aUuidType, "uuidver", FromEnvP("UUIDVER", "4").(string), sUuidType)
	flag.StringVar(&aUuidType, "uv", FromEnvP("UUIDVER", "4").(string), sUuidType)
	flag.BoolVar(&aVersion, "version", FromEnvP("VERSION", false).(bool), sVersion)
	flag.BoolVar(&aVersion, "v", FromEnvP("VERSION", false).(bool), sVersion)
	flag.BoolVar(&aXtra, "xtra", FromEnvP("XTRA", false).(bool), sXtra)
	flag.BoolVar(&aXtra, "x", FromEnvP("XTRA", false).(bool), sXtra)
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

	enc := base58Encoder{}
	var (
		err error
		luu uuid.UUID
		suu string
		u64 uint64
		u32 uint32
	)

	// Lengthen or shorten an existing UUID or generate a new one
	if aUuid != "" {
		luu, err = uuid.Parse(aUuid)

		// It might be a short UUID already
		if err != nil {
			luu, err = enc.Decode(aUuid)
		}
	} else {
		// A non-empty name but default version means we probably want UUIDv5
		if aName != "" && aUuidType == "4" {
			aUuidType = "5"
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

		switch strings.ToUpper(aUuidType) {
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

	suu = enc.Encode(luu)
	if aLong {
		fmt.Println(luu)
	} else {
		fmt.Println(suu)
	}
	if aXtra {
		fmt.Println(xtraInfo(luu))
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
//   https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format-03  (check after 2022-10-02)
//   https://datatracker.ietf.org/doc/html/rfc4122
//   https://en.wikipedia.org/wiki/Universally_unique_identifier

// Other alphabets???
//   https://datatracker.ietf.org/doc/html/draft-msporny-base58-03
//   https://stackoverflow.com/questions/41996761/golang-number-base-conversion/48362821#48362821
// base57    '23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
// base58    '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
// base62    '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
// base64    '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/'
// base64url '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_'

// Other features???
//   https://github.com/yeqown/go-qrcode  generating a barcode bitmap
//   https://github.com/signintech/gopdf  generating a PDF

func Genv1() (uuid.UUID, error) {
	luu, err := uuid.NewUUID()
	return luu, err
}

func Genv2(domain string, id uint32) (uuid.UUID, error) {
	// uuid.SetNodeID([]byte{00, 00, 00, 00, 00, 01})
	switch strings.ToLower(domain) {
	case "person":
		return uuid.NewDCESecurity(uuid.Person, id)
	case "group":
		return uuid.NewDCESecurity(uuid.Group, id)
	case "org":
		return uuid.NewDCESecurity(uuid.Org, id)
	default:
		return uuid.Nil, errors.New("Unsupported domain")
	}
}

func Genv3(name string, namespace string) (uuid.UUID, error) {
	switch strings.ToUpper(namespace) {
	case "DNS":
		return uuid.NewMD5(uuid.NameSpaceDNS, []byte(name)), nil
	case "OID":
		return uuid.NewMD5(uuid.NameSpaceOID, []byte(name)), nil
	case "URL":
		return uuid.NewMD5(uuid.NameSpaceURL, []byte(name)), nil
	case "X500":
		return uuid.NewMD5(uuid.NameSpaceX500, []byte(name)), nil
	default:
		return uuid.Nil, errors.New("Unsupported namespace")
	}
}

func Genv4() (uuid.UUID, error) {
	luu, err := uuid.NewRandom()
	return luu, err
}

func Genv5(name string, namespace string) (uuid.UUID, error) {
	switch strings.ToUpper(namespace) {
	case "DNS":
		return uuid.NewSHA1(uuid.NameSpaceDNS, []byte(name)), nil
	case "OID":
		return uuid.NewSHA1(uuid.NameSpaceOID, []byte(name)), nil
	case "URL":
		return uuid.NewSHA1(uuid.NameSpaceURL, []byte(name)), nil
	case "X500":
		return uuid.NewSHA1(uuid.NameSpaceX500, []byte(name)), nil
	default:
		return uuid.Nil, errors.New("Unsupported namespace")
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
