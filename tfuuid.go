package main

import (
	"fmt"
	// "os"

	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid"
)

type base58Encoder struct{}

func (enc base58Encoder) Encode(u uuid.UUID) string {
	return base58.Encode(u[:])
}

func (enc base58Encoder) Decode(s string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(s))
}

func main() {
	// default alphabet '23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz' (base57)
	// desired alphabet '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz' (base58)

	// Generate a new short UUIDv4
	enc := base58Encoder{}
	id := shortuuid.NewWithEncoder(enc)

	// Encode (shorten) an existing UUID
	// uu, err := uuid.Parse("8966e5ee-445e-401b-a921-bf5020c516d2")
	// if err != nil {
	//   os.Exit(1)
	// }}
	// enc := base58Encoder{}
	// id := enc.Encode(uu)
	// Hy5v2PaRjQhVB172zb6fpD

	// Decode (lengthen) an existing shortUUID
	// enc := base58Encoder{}
	// id, err := enc.Decode("Hy5v2PaRjQhVB172zb6fpD")
	// if err != nil {
	//   os.Exit(1)
	// }}
	// 8966e5ee-445e-401b-a921-bf5020c516d2

	// Generate a UUIDv5
	// name := []byte("python.org")
	// uu := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(name))
	// uu := uuid.NewSHA1(uuid.NameSpaceOID, []byte(name))
	// uu := uuid.NewSHA1(uuid.NameSpaceURL, []byte(name))
	// uu := uuid.NewSHA1(uuid.NameSpaceX500, []byte(name))
	// enc := base58Encoder{}
	// id := enc.Encode(uu)

	fmt.Println(id)
}

// https://github.com/lithammer/shortuuid  go implementation
// https://pkg.go.dev/github.com/lithammer/shortuuid  go implementation
// https://github.com/skorokithakis/shortuuid  python implementation that is compatible
// https://pypi.org/project/shortuuid/  python implementation that is compatible
// https://github.com/yeqown/go-qrcode  generating a barcode bitmap
// https://github.com/signintech/gopdf  generating a PDF
