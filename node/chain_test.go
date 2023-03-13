package node

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/elyutikov/goblockchain/crypto"
	"github.com/elyutikov/goblockchain/proto"
	"github.com/elyutikov/goblockchain/types"
	"github.com/elyutikov/goblockchain/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func randomBlock(t *testing.T, chain *Chain) *proto.Block {
	privKey := crypto.GeneratePrivateKey()
	b := util.RandomBlock()
	prevBlock, err := chain.GetBlockByHeight(chain.Height())
	require.Nil(t, err)
	b.Header.PrevHash = types.HashBlock(prevBlock)
	types.SignBlock(privKey, b)
	return b
}

func TestNewChain(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore(), NewMemoryTxStore(), NewMemoryUTXOStore())
	assert.Equal(t, 0, chain.Height())
	_, err := chain.GetBlockByHeight(0)
	assert.Nil(t, err)
}

func TestChainHeight(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore(), NewMemoryTxStore(), NewMemoryUTXOStore())

	for i := 0; i < 100; i++ {
		b := randomBlock(t, chain)
		assert.Nil(t, chain.AddBlock(b))
		assert.Equal(t, chain.Height(), i+1)
	}
}

func TestAddBlock(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore(), NewMemoryTxStore(), NewMemoryUTXOStore())

	for i := 0; i < 100; i++ {
		block := randomBlock(t, chain)
		blockHash := types.HashBlock(block)

		assert.Nil(t, chain.AddBlock(block))

		fetchedBlockByHash, err := chain.GetBlockByHash(blockHash)
		assert.Nil(t, err)
		assert.Equal(t, fetchedBlockByHash, block)

		fetchedBlockByHeight, err := chain.GetBlockByHeight(i + 1)
		assert.Nil(t, err)
		assert.Equal(t, fetchedBlockByHeight, block)
	}
}

func TestAddBlockWithTx(t *testing.T) {
	var (
		chain     = NewChain(NewMemoryBlockStore(), NewMemoryTxStore(), NewMemoryUTXOStore())
		block     = randomBlock(t, chain)
		privKey   = crypto.NewPrivateKeyFromSeedStr(goodSeed)
		recipient = privKey.Public().Address().Bytes()
	)

	ftt, err := chain.txStore.Get("43636e981159793eccd37f1a32a287752a7cd5c6faa16b76d25374308081ffd0")
	assert.Nil(t, err)

	inputs := []*proto.TxInput{
		{
			PrevTxHash:   types.HashTransaction(ftt),
			PrevOutIndex: 0,
			PublicKey:    privKey.Public().Bytes(),
		},
	}
	outputs := []*proto.TxOutput{
		{
			Amount:  100,
			Address: recipient,
		},
		{
			Amount:  900,
			Address: privKey.Public().Address().Bytes(),
		},
	}
	tx := &proto.Transaction{
		Version: 1,
		Inputs:  inputs,
		Outputs: outputs,
	}

	sig := types.SignTransaction(privKey, tx)
	tx.Inputs[0].Signature = sig.Bytes()

	block.Transactions = append(block.Transactions, tx)
	require.Nil(t, chain.AddBlock(block))
	txHash := hex.EncodeToString(types.HashTransaction(tx))

	fetchedTx, err := chain.txStore.Get(txHash)
	assert.Nil(t, err)
	assert.Equal(t, tx, fetchedTx)

	address := crypto.AddressFromBytes(tx.Outputs[0].Address)
	key := fmt.Sprintf("%s_%s", address, txHash)
	utxo, err := chain.utxoStore.Get(key)
	assert.Nil(t, err)

	fmt.Println(utxo)
}
