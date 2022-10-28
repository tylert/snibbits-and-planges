package main

import (
	"fmt"
	// "log"
	"os"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/google/uuid"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

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
			luu, err = Genv1(aNodeId)
		case "2":
			luu, err = Genv2(aNodeId, aDomain, u32)
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
	switch strings.ToLower(aEncoding) {
	case "base58":
		fmt.Println(suu)
	case "none":
		fmt.Println(luu)
	default:
		fmt.Println("Unrecognized encoding")
		os.Exit(4)
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
//   https://pkg.go.dev/github.com/google/uuid
//   https://github.com/google/uuid
//   https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format-04
//   http://gh.peabody.io/uuidv6/
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

func xtraInfo(luu uuid.UUID) string {
	ver := strings.Split(luu.Version().String(), "_")
	output := fmt.Sprintf("UUID Version:%s Variant:%s", ver[1], luu.Variant())

	switch ver[1] {
	case "2":
		output += fmt.Sprintf(" Domain:%s Id:%d", luu.Domain().String(), luu.ID())
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
