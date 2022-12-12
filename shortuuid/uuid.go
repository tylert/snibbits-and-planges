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
//   https://pkg.go.dev/github.com/google/uuid
//   https://github.com/google/uuid
//   https://github.com/uuid6/uuid6-ietf-draft
//   https://github.com/ietf-wg-uuidrev/rfc4122bis !!!
//   https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format-04
//   http://gh.peabody.io/uuidv6/
//   https://datatracker.ietf.org/doc/html/rfc4122
//   https://en.wikipedia.org/wiki/Universally_unique_identifier

func GenUUIDv1(nodeid string) (string, error) {
	// uuid.SetClockSequence(-1)
	if strings.ToUpper(nodeid) == "RANDOM" {
		b := make([]byte, 6)
		_, err := rand.Read(b)
		if err != nil {
			panic("random fail")
		}
		uuid.SetNodeID(b)
	} else {
		uuid.SetNodeInterface(nodeid)
	}

	uu, err := uuid.NewUUID()
	return uu.String(), err
}

func GenUUIDv2(nodeid string, domain string, id uint32) (string, error) {
	// uuid.SetClockSequence(-1)
	if strings.ToUpper(nodeid) == "RANDOM" {
		b := make([]byte, 6)
		_, err := rand.Read(b)
		if err != nil {
			panic("random fail")
		}
		uuid.SetNodeID(b)
	} else {
		uuid.SetNodeInterface(nodeid)
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

func GenUUIDv3(name string, namespace string) (string, error) {
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

func GenUUIDv5(name string, namespace string) (string, error) {
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

// func GenUUIDv6(nodeid string) (string, error) {
// func GenUUIDv7() (string, error) {
// func GenUUIDv8() (string, error) {
