package main

import (
	"errors"
	"flag"
	"fmt"
	// "log"
	"os"
	"strings"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/google/uuid"
	// "github.com/nicksnyder/basen"
)

// func Genv1() (uuid.UUID, error) {
// }

// func Genv2(domain string, id uint32) (uuid.UUID, error) {
//   switch domain {
//   case "Person":
//     return uuid.NewDCESecurity(uuid.Person, id)
//   case "Group":
//     return uuid.NewDCESecurity(uuid.Group, id)
//   case "Org":
//     return uuid.NewDCESecurity(uuid.Org, id)
//   }
// }

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
	// speed up batch operations at the cost of a bit more danger???
	// uuid.EnableRandPool()
	// uuid.DisableRandPool()

	lluu, err := uuid.NewRandom()
	return lluu, err
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

func (enc base58Encoder) Encode(lluu uuid.UUID) string {
	return base58.Encode(lluu[:])
}

func (enc base58Encoder) Decode(shuu string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(shuu))
}

var (
	n  = flag.String("name", "", "Name")
	ns = flag.String("namespace", "DNS", "Namespace")
)

func init() {
	flag.StringVar(n, "n", "", "Name")
	flag.StringVar(ns, "ns", "DNS", "Namespace")
}

func main() {
	// base57    '23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
	// base58    '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
	// base62    '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
	// base64    '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/'
	// base64url '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_'

	flag.Parse()

	// lluu, err := Genv3("python.org", "URL")
	// lluu, err := Genv4()
	lluu, err := Genv5(*n, *ns)
	// lluu, err := uuid.Parse("cd5d0bff-2444-5d26-ab53-4f7db1cb733d")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	enc := base58Encoder{}
	// lluu, err := enc.Decode("SMqCfPLDiH5aTTgLmGR4np")
	shuu := enc.Encode(lluu)

	// fmt.Println(lluu)
	fmt.Println(shuu)
}

// https://stackoverflow.com/questions/41996761/golang-number-base-conversion/48362821#48362821

// https://github.com/yeqown/go-qrcode  generating a barcode bitmap
// https://github.com/signintech/gopdf  generating a PDF
// https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format-03  (check after 2022-10-02)
// https://datatracker.ietf.org/doc/html/draft-msporny-base58-03
