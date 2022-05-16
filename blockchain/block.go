package blockchain

type Block struct {
	PreviosHash []byte
	Hash        []byte
	Data        []byte
	Nouce       int
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

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
