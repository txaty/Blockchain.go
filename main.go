package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain is an array of pointers to Blocks
type BlockChain struct {
	blocks []*Block
}

// Block is the block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash creates a new hash based on the data and the previous hash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) // Simpler than the real hash algorithm for the blockchain
	b.Hash = hash[:]
}

// CreateBlock creates a new Block by Block constructor, and then calculate the hash of it
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock creates a new block and append it to the chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Genesis creates a new block containing data "Genesis"
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain initialize a Blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x \n", block.PrevHash)
		s := string(block.Data)
		fmt.Printf("Data in Block: %s \n", s)
		fmt.Printf("Hash: %x \n", block.Hash)
	}
}
