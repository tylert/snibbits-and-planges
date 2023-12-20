package main

import (
	"crypto/rand"
	"errors"
	"strings"

	"github.com/google/uuid"
)

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

func GenUUIDv6(node string) (string, error) {
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

	uu, err := uuid.NewV6()
	return uu.String(), err
}

func GenUUIDv7() (string, error) {
	uu, err := uuid.NewV7()
	return uu.String(), err
}
