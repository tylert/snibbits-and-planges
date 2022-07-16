package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/sethvargo/go-diceware/diceware"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		log.Fatal(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatal(err)
	}
	secret_list, err := diceware.Generate(24)
	if err != nil {
		log.Fatal(err)
	}
	secret := strings.Join(secret_list, " ")
	// seed := bip39.NewSeed(mnemonic, secret)

	fmt.Println(mnemonic)
	fmt.Println(secret)
}
