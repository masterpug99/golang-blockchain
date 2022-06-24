package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Make a BlockChain
type BlockChain struct {
	blocks []*Block
}

// Make a Block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Make a Block's Hash
//
// if the previousBlock's Info is changed the blocks' Hash following this
// Block are all changed
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Create Block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Add a Block to BlockChain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Make a Genesis Bloc
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Make a Blockchain with Genesis Block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second21 Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.blocks {
		//fmt.Printf("Previous Hash : %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
