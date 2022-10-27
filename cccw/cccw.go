package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sethvargo/go-diceware/diceware"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// Print out the version information
	if aVersion {
		fmt.Println(GetVersion())
		os.Exit(0)
	}

	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Fatal(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatal(err)
	}
	secret_list, err := diceware.Generate(12)
	if err != nil {
		log.Fatal(err)
	}
	secret := strings.Join(secret_list, " ")
	// seed := bip39.NewSeed(mnemonic, secret)

	fmt.Println(mnemonic)
	fmt.Println(secret)
}
