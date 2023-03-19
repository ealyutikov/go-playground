package node

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/elyutikov/goblockchain/proto"
	"github.com/elyutikov/goblockchain/types"
)

type BlockStore interface {
	Put(b *proto.Block) error
	Get(hash string) (*proto.Block, error)
}

type MemoryBlockStore struct {
	lock   sync.RWMutex
	blocks map[string]*proto.Block
}

func NewMemoryBlockStore() *MemoryBlockStore {
	return &MemoryBlockStore{
		blocks: make(map[string]*proto.Block),
	}
}

func (s *MemoryBlockStore) Put(b *proto.Block) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	hash := hex.EncodeToString(types.HashBlock(b))
	s.blocks[hash] = b
	return nil
}

func (s *MemoryBlockStore) Get(hash string) (*proto.Block, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	block, ok := s.blocks[hash]
	if !ok {
		return nil, fmt.Errorf("block with hash [%s] does not exists", hash)
	}
	return block, nil
}
