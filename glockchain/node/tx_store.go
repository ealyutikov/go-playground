package node

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/elyutikov/goblockchain/proto"
	"github.com/elyutikov/goblockchain/types"
)

type TxStore interface {
	Put(transaction *proto.Transaction) error
	Get(string) (*proto.Transaction, error)
}

type MemoryTxStore struct {
	lock sync.RWMutex
	txx  map[string]*proto.Transaction
}

func NewMemoryTxStore() *MemoryTxStore {
	return &MemoryTxStore{
		txx: make(map[string]*proto.Transaction),
	}
}

func (store *MemoryTxStore) Put(tx *proto.Transaction) error {
	store.lock.Lock()
	defer store.lock.Unlock()

	hash := hex.EncodeToString(types.HashTransaction(tx))
	store.txx[hash] = tx
	return nil
}

func (store *MemoryTxStore) Get(hash string) (*proto.Transaction, error) {
	store.lock.RLock()
	defer store.lock.RUnlock()

	tx, ok := store.txx[hash]
	if !ok {
		return nil, fmt.Errorf("unanble to find tx with hash [%s]", hash)
	}
	return tx, nil
}
