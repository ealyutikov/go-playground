package types

import (
	"testing"

	"github.com/elyutikov/goblockchain/crypto"
	"github.com/elyutikov/goblockchain/proto"
	"github.com/elyutikov/goblockchain/util"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	var (
		fromPrivKey = crypto.GeneratePrivateKey()
		fromAddress = fromPrivKey.Public().Address().Bytes()

		toPrivKey = crypto.GeneratePrivateKey()
		toAddress = toPrivKey.Public().Address().Bytes()
	)

	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}

	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}

	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Verrsion: 1,
		Inputs:   []*proto.TxInput{input},
		Outputs:  []*proto.TxOutput{output1, output2},
	}

	sig := SignTransaction(fromPrivKey, tx)
	input.Signature = sig.Bytes()

	assert.True(t, VerifyTransaction(tx))
}
