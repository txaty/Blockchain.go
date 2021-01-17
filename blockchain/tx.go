package blockchain

import (
	"blockchain_test/wallet"
	"bytes"
)

// TxOutput records transaction output
type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

// TxInput records transaction input
type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

// NewTXOutput creates a new transaction and locks it
func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

// UsesKey compares the locking hash with public key hash
func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

// Lock locks the output
func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// IsLockedWithKey checks if the output has been locked with a key
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}
