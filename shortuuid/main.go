package main

import (
	"fmt"
	// "log"
	"os"
	"strconv"
	"strings"

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
			enc := base58Encoder{}
			auu, err = enc.Decode(aUuid)
		}
		luu = auu.String()
	} else {
		// A non-empty name but default type means we probably want UUIDv5
		if aName != "" && aType == "4" {
			aType = "5"
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

		switch strings.ToUpper(aType) {
		case "1":
			luu, err = GenUUIDv1(aNodeId)
		case "2":
			luu, err = GenUUIDv2(aNodeId, aDomain, u32)
		case "3":
			luu, err = GenUUIDv3(aName, aNamespace)
		case "4":
			luu, err = GenUUIDv4()
		case "5":
			luu, err = GenUUIDv5(aName, aNamespace)
		// case "6":
		// 	luu, err = GenUUIDv6()
		// case "7":
		// 	luu, err = GenUUIDv7()
		// case "8":
		// 	luu, err = GenUUIDv8()
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
	enc := base58Encoder{}
	suu = enc.Encode(auu)
	switch strings.ToUpper(aEncoding) {
	case "BASE58":
		fmt.Println(suu)
	case "NONE":
		fmt.Println(luu)
	default:
		fmt.Println("Unsupported encoding")
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

// Other features???
//   https://paulgorman.org/technical/blog/20171113164018.html  JSON config
//   https://github.com/signintech/gopdf  generating a PDF

func xtraInfo(luu uuid.UUID) string {
	ver := strings.Split(luu.Version().String(), "_")
	output := fmt.Sprintf("UUIDv%s Variant:%s", ver[1], luu.Variant())

	switch ver[1] {
	case "2":
		output += fmt.Sprintf(" Domain:%s Id:%d", luu.Domain().String(), luu.ID())
	}

	return output
}
