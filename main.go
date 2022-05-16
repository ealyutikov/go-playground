package main

import (
	"crypto/sha256"
	"fmt"

	"bytes"

	"rsc.io/quote"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	PreviosHash []byte
	Hash        []byte
	Data        []byte
}

func (b *Block) calculateHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviosHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		PreviosHash: prevHash,
		Hash:        []byte{},
		Data:        []byte(data),
	}
	block.calculateHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func IniteBlockChain() *BlockChain {
	return &BlockChain{
		blocks: []*Block{Genesis()},
	}
}

func main() {
	chain := IniteBlockChain()

	chain.AddBlock("first block")
	chain.AddBlock("second block")
	chain.AddBlock("third block")

	for _, block := range chain.blocks {
		fmt.Printf("Previos hash: %x\n", block.PreviosHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Println()
	}

	fmt.Println(quote.Hello())
}
