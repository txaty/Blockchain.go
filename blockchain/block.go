package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain is an array of pointers to Blocks
type BlockChain struct {
	Blocks []*Block
}

// Block is the block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// DeriveHash creates a new hash based on the data and the previous hash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) // Simpler than the real hash algorithm for the blockchain
	b.Hash = hash[:]
}

// CreateBlock creates a new Block by Block constructor, and then calculate the hash of it
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock creates a new block and append it to the chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis creates a new block containing data "Genesis"
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain initialize a Blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
