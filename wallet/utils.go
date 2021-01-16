package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

// Base58Encode encodes base58
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

// Base58Decode decodes base58
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}

// 0 O l I + /
