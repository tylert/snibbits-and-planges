package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/simukti/totp"
)

func main() {
	// Print out the version information
	if aVersion {
		fmt.Println(GetVersion())
		os.Exit(0)
	}

	var (
		err error
		ttp *totp.TOTP
	)

	switch strings.ToUpper(aType) {
	case "SHA1":
		ttp, err = totp.NewSHA1(aDigits, aCount)
	case "SHA256":
		ttp, err = totp.NewSHA256(aDigits, aCount)
	case "SHA512":
		ttp, err = totp.NewSHA512(aDigits, aCount)
	default:
		fmt.Println("Unsupported hash algorithm")
		os.Exit(1)
	}

	if err != nil {
		panic("Got error making hash!!!")
	}

	// otp, err := ttp.OTP([]byte(aSeed), time.Unix(111111111, 0))
	otp, err := ttp.OTP([]byte(aSeed), time.Now())
	if err != nil {
		panic("Got error making token!!!")
	}

	fmt.Println(otp)
}
