package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
)

// NewPriPub ...
func NewPriPub() (string, string) {
	pri, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", ""
	}
	return hex.EncodeToString(pri.Serialize()), hex.EncodeToString(pri.PubKey().SerializeCompressed())
}

func main() {
	pri, pub := NewPriPub()
	fmt.Println("ecc pri:", pri)
	fmt.Println("ecc pub:", pub)
}
