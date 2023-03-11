package node

import (
	"encoding/hex"

	"github.com/elyutikov/goblockchain/proto"
	"github.com/elyutikov/goblockchain/types"
)

type Mempool struct {
	txx map[string]*proto.Transaction
}

func NewMempool() *Mempool {
	return &Mempool{txx: make(map[string]*proto.Transaction)}
}

func (pool *Mempool) Has(tx *proto.Transaction) bool {
	hash := hex.EncodeToString(types.HashTransaction(tx))
	tx, ok := pool.txx[hash]
	return ok
}

func (pool *Mempool) Add(tx *proto.Transaction) bool {
	if pool.Has(tx) {
		return false
	}

	hash := hex.EncodeToString(types.HashTransaction(tx))
	pool.txx[hash] = tx
	return true
}
