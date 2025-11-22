package main

import (
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/google/uuid"
	// "github.com/tv42/base58"
	// "github.com/tv42/zbase32"
)

// Other encodings???
//   https://devtoolsdaily.com/blog/short-uuid
//   https://github.com/uuid25/go-uuid25
//   https://stackoverflow.com/questions/37934162/output-uuid-in-go-as-a-short-string
//   https://stackoverflow.com/questions/41996761/golang-number-base-conversion/48362821#48362821
//   https://datatracker.ietf.org/doc/html/draft-msporny-base58-03
//   https://github.com/Dasio/base45/blob/main/base45.go
//   https://cs.opensource.google/go/go/+/master:src/encoding/base64/base64.go
//   https://rfc-editor.org/info/rfc9285
// base57    '23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
// base58    '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
// base62    '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'
// base64    '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/'
// base64url '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_'

type base58Encoder struct{}

func (enc base58Encoder) EncodeBase58(luu uuid.UUID) string {
	return base58.Encode(luu[:])
}

func (enc base58Encoder) DecodeBase58(suu string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(suu))
}
