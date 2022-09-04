package main

import (
	"errors"
	"flag"
	"fmt"
	// "log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/google/uuid"
	// "github.com/nicksnyder/basen"
)

func Genv3(n string, ns string) (uuid.UUID, error) {
	switch strings.ToUpper(ns) {
	case "DNS":
		return uuid.NewMD5(uuid.NameSpaceDNS, []byte(n)), nil
	case "OID":
		return uuid.NewMD5(uuid.NameSpaceOID, []byte(n)), nil
	case "URL":
		return uuid.NewMD5(uuid.NameSpaceURL, []byte(n)), nil
	case "X500":
		return uuid.NewMD5(uuid.NameSpaceX500, []byte(n)), nil
	default:
		return uuid.Nil, errors.New("Invalid namespace")
	}
}

func Genv4() (uuid.UUID, error) {
	luu, err := uuid.NewRandom()
	return luu, err
}

func Genv5(n string, ns string) (uuid.UUID, error) {
	switch strings.ToUpper(ns) {
	case "DNS":
		return uuid.NewSHA1(uuid.NameSpaceDNS, []byte(n)), nil
	case "OID":
		return uuid.NewSHA1(uuid.NameSpaceOID, []byte(n)), nil
	case "URL":
		return uuid.NewSHA1(uuid.NameSpaceURL, []byte(n)), nil
	case "X500":
		return uuid.NewSHA1(uuid.NameSpaceX500, []byte(n)), nil
	default:
		return uuid.Nil, errors.New("Invalid namespace")
	}
}

// func Genv6() (uuid.UUID, error) {
// func Genv7() (uuid.UUID, error) {
// func Genv8() (uuid.UUID, error) {

type base58Encoder struct{}

func (enc base58Encoder) Encode(luu uuid.UUID) string {
	return base58.Encode(luu[:])
}

func (enc base58Encoder) Decode(suu string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(suu))
}

// Long options
var (
	l  = flag.Bool("long", false, "Show the long UUID instead of the short one")
	n  = flag.String("name", "", "Name to use for UUIDv5 or v3 hash")
	ns = flag.String("namespace", "DNS", "Namespace to use for UUIDv5 or v3 hash")
	u  = flag.String("uuid", "", "Existing UUID to shorten or lengthen")
	uv = flag.String("uuidver", "4", "Generate UUIDv5, v4 or v3")
	v  = flag.Bool("version", false, "Display version information and exit")
)

// Short options
func init() {
	flag.BoolVar(l, "l", false, "Show the long UUID instead of the short one")
	flag.StringVar(n, "n", "", "Name to use for UUIDv5 or v3 hash")
	flag.StringVar(ns, "ns", "DNS", "Namespace to use for UUIDv5 or v3 hash")
	flag.StringVar(u, "u", "", "Existing UUID to shorten or lengthen")
	flag.StringVar(uv, "uv", "4", "Generate UUIDv5, v4 or v3")
	flag.BoolVar(v, "v", false, "Display version information and exit")
}

// go build -ldflags "-X main.Version=$(git describe --always --dirty --tags)"
var Version string = ""

func main() {
	flag.Parse()

	// Print out the version information
	if *v {
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
		os.Exit(0)
	}

	enc := base58Encoder{}
	var (
		luu uuid.UUID
		suu string
		err error
	)

	// Lengthen or shorten an existing UUID or generate a new one
	if *u != "" {
		luu, err = uuid.Parse(*u)

		// It might be a short UUID already
		if err != nil {
			luu, err = enc.Decode(*u)
		}
	} else {
		// A non-empty name but default version means we probably want UUIDv5
		if *n != "" && *uv == "4" {
			*uv = "5"
		}

		switch strings.ToUpper(*uv) {
		case "3":
			luu, err = Genv3(*n, *ns)
		case "4":
			luu, err = Genv4()
		case "5":
			luu, err = Genv5(*n, *ns)
		default:
			fmt.Println("Invalid UUID version")
			os.Exit(2)
		}
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	suu = enc.Encode(luu)
	if *l {
		fmt.Println(luu)
	} else {
		fmt.Println(suu)
	}
}

// base57    '23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
// base58    '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
// base62    '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
// base64    '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/'
// base64url '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_'

// https://stackoverflow.com/questions/41996761/golang-number-base-conversion/48362821#48362821
// https://github.com/yeqown/go-qrcode  generating a barcode bitmap
// https://github.com/signintech/gopdf  generating a PDF
// https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format-03  (check after 2022-10-02)
// https://datatracker.ietf.org/doc/html/draft-msporny-base58-03
