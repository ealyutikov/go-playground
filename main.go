package main

import (
	"fmt"
	"go-blockchain/blockchain"
	"strconv"
)

func main() {
	chain := blockchain.IniteBlockChain()

	chain.AddBlock("first block")
	chain.AddBlock("second block")
	chain.AddBlock("third block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PreviosHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Println()

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
