package node

import (
	"fmt"
	"sync"
)

type UTXO struct {
	Hash     string
	OutIndex int
	Amount   int64
	Spent    bool
}

type UTXOStore interface {
	Put(string, *UTXO) error
	Get(string) (*UTXO, error)
}

type MemoryUTXOStore struct {
	lock sync.RWMutex
	data map[string]*UTXO
}

func NewMemoryUTXOStore() *MemoryUTXOStore {
	return &MemoryUTXOStore{
		data: make(map[string]*UTXO),
	}
}

func (store *MemoryUTXOStore) Put(key string, utxo *UTXO) error {
	store.lock.Lock()
	defer store.lock.Unlock()

	store.data[key] = utxo
	return nil
}

func (store *MemoryUTXOStore) Get(hash string) (*UTXO, error) {
	store.lock.RLock()
	defer store.lock.RUnlock()

	utxo, ok := store.data[hash]
	if !ok {
		return nil, fmt.Errorf("unable to find utxo with hash %s", hash)
	}

	return utxo, nil
}
