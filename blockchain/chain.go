package blockchain

type BlockChain struct {
	Blocks []*Block
}

func IniteBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{Genesis()},
	}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	block := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, block)
}
