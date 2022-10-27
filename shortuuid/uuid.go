package main

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	// uuid6 "github.com/bradleypeabody/gouuidv6"
	// uuid7 "github.com/uuid6/uuid6go-proto"
)

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
