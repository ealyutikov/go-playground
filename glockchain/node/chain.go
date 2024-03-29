package node

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/elyutikov/goblockchain/crypto"
	"github.com/elyutikov/goblockchain/proto"
	"github.com/elyutikov/goblockchain/types"
)

const goodSeed = "535a86e30920723b4f0f176d2f69456a20bba38aa179675108be3971b9736fc3"

type Chain struct {
	blockStore BlockStore
	txStore    TxStore
	utxoStore  UTXOStore
	headers    *HeaderList
}

func NewChain(bs BlockStore, tx TxStore, utxo UTXOStore) *Chain {
	chain := &Chain{
		blockStore: bs,
		txStore:    tx,
		utxoStore:  utxo,
		headers:    NewHeaderList(),
	}
	chain.addBlock(createGenesisBlock())
	return chain
}

func (c *Chain) Height() int {
	return c.headers.Height()
}

func (c *Chain) AddBlock(b *proto.Block) error {
	if err := c.ValidateBlock(b); err != nil {
		return err
	}
	return c.addBlock(b)
}

func (c *Chain) addBlock(b *proto.Block) error {
	c.headers.Add(b.Header)

	for _, tx := range b.Transactions {
		if err := c.txStore.Put(tx); err != nil {
			return err
		}
		hash := hex.EncodeToString(types.HashTransaction(tx))

		for it, output := range tx.Outputs {
			utxo := &UTXO{
				Hash:     hash,
				OutIndex: it,
				Amount:   output.Amount,
				Spent:    false,
			}
			if err := c.utxoStore.Put(utxo); err != nil {
				return err
			}
		}
	}

	return c.blockStore.Put(b)
}

func (c *Chain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashHex := hex.EncodeToString(hash)
	return c.blockStore.Get(hashHex)
}

func (c *Chain) GetBlockByHeight(height int) (*proto.Block, error) {
	if height > c.Height() {
		return nil, fmt.Errorf("given height [%d] is too height [%d]", height, c.Height())
	}
	header := c.headers.Get(height)
	hash := types.HashHeader(header)
	return c.GetBlockByHash(hash)
}

func (c *Chain) ValidateBlock(b *proto.Block) error {
	// Validate the block's signature
	if !types.VerifyBlock(b) {
		return fmt.Errorf("invalid block signature")
	}

	// Validate the previous hash
	currentBlock, err := c.GetBlockByHeight(c.Height())
	if err != nil {
		return err
	}
	hash := types.HashBlock(currentBlock)
	if !bytes.Equal(hash, b.Header.PrevHash) {
		return fmt.Errorf("invalid previos block hash")
	}

	for _, tx := range b.Transactions {
		if err := c.ValidateTransaction(tx); err != nil {
			return err
		}
	}

	return nil
}

func (c *Chain) ValidateTransaction(tx *proto.Transaction) error {
	// verify the signature
	if !types.VerifyTransaction(tx) {
		return fmt.Errorf("invalid tx signature")
	}

	// verify if all outputs are unspent
	hash := hex.EncodeToString(types.HashTransaction(tx))
	sumInputs := 0
	for i := 0; i < len(tx.Inputs); i++ {
		prevHash := hex.EncodeToString(tx.Inputs[i].PrevTxHash)
		key := fmt.Sprintf("%s_%d", prevHash, i)
		utxo, err := c.utxoStore.Get(key)
		sumInputs += int(utxo.Amount)
		if err != nil {
			return err
		}
		if utxo.Spent {
			return fmt.Errorf("input %d of tx %s is already spent", i, hash)
		}
	}

	sumOutputs := 0
	for _, output := range tx.Outputs {
		sumOutputs += int(output.Amount)
	}

	if sumInputs < sumOutputs {
		return fmt.Errorf("insufficient balance got (%d) spending (%d)", sumInputs, sumOutputs)
	}

	return nil
}

func createGenesisBlock() *proto.Block {
	privKey := crypto.NewPrivateKeyFromSeedStr(goodSeed)
	block := &proto.Block{
		Header: &proto.Header{
			Version: 1,
		},
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{},
		Outputs: []*proto.TxOutput{
			{
				Amount:  1000,
				Address: privKey.Public().Address().Bytes(),
			},
		},
	}
	block.Transactions = append(block.Transactions, tx)
	types.SignBlock(privKey, block)
	return block
}
