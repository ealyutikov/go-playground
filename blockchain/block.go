package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

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

	block.Hash = hash
	block.Nouce = nouce

	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
