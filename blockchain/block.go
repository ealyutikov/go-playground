package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	PreviosHash []byte
	Hash        []byte
	Data        []byte
	Nouce       int
}

func (b *Block) CalculateHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviosHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		PreviosHash: prevHash,
		Hash:        []byte{},
		Data:        []byte(data),
		Nouce:       0,
	}

	pow := NewProof(block)
	nouce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nouce = nouce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func IniteBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{Genesis()},
	}
}
