package main

import (
	"crypto/rand"
	"errors"
	"strings"

	"github.com/google/uuid"
	// uuid6 "github.com/bradleypeabody/gouuidv6"
	// uuid7 "github.com/uuid6/uuid6go-proto"
	// uuid8 ???
)

// Other UUID stuff???
//   https://pkg.go.dev/github.com/pborman/uuid
//   https://github.com/pborman/uuid
//   https://pkg.go.dev/github.com/google/uuid
//   https://github.com/google/uuid
//   https://github.com/uuid6/uuid6-ietf-draft
//   https://github.com/ietf-wg-uuidrev/rfc4122bis
//   https://www.ietf.org/archive/id/draft-ietf-uuidrev-rfc4122bis-07.html
//   http://gh.peabody.io/uuidv6/
//   https://datatracker.ietf.org/doc/html/rfc4122
//   https://en.wikipedia.org/wikiwiki/Universally_unique_identifier

func GenUUIDv1(node string) (string, error) {
	// uuid.SetClockSequence(-1)
	if strings.ToUpper(node) == "RANDOM" {
		b := make([]byte, 6)
		_, err := rand.Read(b)
		if err != nil {
			panic("random fail")
		}
		uuid.SetNodeID(b)
	} else {
		uuid.SetNodeInterface(node)
	}

	uu, err := uuid.NewUUID()
	return uu.String(), err
}

func GenUUIDv2(node string, domain string, id uint32) (string, error) {
	// uuid.SetClockSequence(-1)
	if strings.ToUpper(node) == "RANDOM" {
		b := make([]byte, 6)
		_, err := rand.Read(b)
		if err != nil {
			panic("random fail")
		}
		uuid.SetNodeID(b)
	} else {
		uuid.SetNodeInterface(node)
	}

	switch strings.ToUpper(domain) {
	case "PERSON":
		uu, err := uuid.NewDCESecurity(uuid.Person, id)
		return uu.String(), err
	case "GROUP":
		uu, err := uuid.NewDCESecurity(uuid.Group, id)
		return uu.String(), err
	case "ORG":
		uu, err := uuid.NewDCESecurity(uuid.Org, id)
		return uu.String(), err
	default:
		return uuid.Nil.String(), errors.New("Unsupported domain")
	}
}

func GenUUIDv3(namespace string, name string) (string, error) {
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

func GenUUIDv4() (string, error) {
	uu, err := uuid.NewRandom()
	return uu.String(), err
}

func GenUUIDv5(namespace string, name string) (string, error) {
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

// func GenUUIDv6(node string) (string, error) {
//   uu, err := GenUUIDv1(node)

// #include <stdio.h>
// #include <stdint.h>
// #include <inttypes.h>
// #include <arpa/inet.h>
// #include <uuid/uuid.h>
//
// /* Converts UUID version 1 to version 6 in place. */
// void uuidv1tov6(uuid_t u) {
//
//   uint64_t ut;
//   unsigned char *up = (unsigned char *)u;
//
//   // load ut with the first 64 bits of the UUID
//   ut = ((uint64_t)ntohl(*((uint32_t*)up))) << 32;
//   ut |= ((uint64_t)ntohl(*((uint32_t*)&up[4])));
//
//   // dance the bit-shift...
//   ut =
//     ((ut >> 32) & 0x0FFF) | // 12 least significant bits
//     (0x6000) | // version number
//     ((ut >> 28) & 0x0000000FFFFF0000) | // next 20 bits
//     ((ut << 20) & 0x000FFFF000000000) | // next 16 bits
//     (ut << 52); // 12 most significant bits
//
//   // store back in UUID
//   *((uint32_t*)up) = htonl((uint32_t)(ut >> 32));
//   *((uint32_t*)&up[4]) = htonl((uint32_t)(ut));
//
// }

// func GenUUIDv7() (string, error) {
// func GenUUIDv8() (string, error) {
