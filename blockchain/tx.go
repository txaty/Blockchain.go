package blockchain

import (
	"blockchain_test/wallet"
	"bytes"
	"encoding/gob"
)

// TxOutput records transaction output
type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

// TxOutputs identifies the transaction output and sort them by unspendable outputs
type TxOutputs struct {
	Outputs []TxOutput
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

// Serialize transaction outputs (TxOutputs) into bytes
func (outs TxOutputs) Serialize() []byte {
	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(outs)
	Handle(err)
	return buffer.Bytes()
}

// DeserializeOutputs deserialize data into TxOutputs
func DeserializeOutputs(data []byte) TxOutputs {
	var outputs TxOutputs
	decode := gob.NewDecoder(bytes.NewReader(data))
	err := decode.Decode(&outputs)
	Handle(err)
	return outputs
}
