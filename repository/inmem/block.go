package inmem

import (
	"context"
	"fmt"
	"sync"

	"github.com/checkaayush/authware/model"
)

var (
	blocks = map[int]*model.Block{
		1: {
			ID:       1,
			Title:    "Block 1",
			MetricID: 1,
		},
	}
	blockSeq  = 2
	blockLock = sync.Mutex{}
)

func (i *inMemRepo) ListBlocks(ctx context.Context) ([]model.Block, error) {
	blockLock.Lock()
	defer blockLock.Unlock()

	bs := []model.Block{}
	for _, block := range blocks {
		bs = append(bs, *block)
	}
	return bs, nil
}

func (i *inMemRepo) GetBlockByID(ctx context.Context, id int) (*model.Block, error) {
	blockLock.Lock()
	defer blockLock.Unlock()

	if block, ok := blocks[id]; ok {
		return block, nil
	}

	return nil, fmt.Errorf("not found")
}

func (i *inMemRepo) CreateBlock(ctx context.Context, b *model.Block) (*model.Block, error) {
	blockLock.Lock()
	defer blockLock.Unlock()

	b.ID = blockSeq
	blocks[b.ID] = b
	blockSeq++
	return b, nil
}

func (i *inMemRepo) DeleteBlock(ctx context.Context, id int) error {
	blockLock.Lock()
	defer blockLock.Unlock()

	delete(blocks, id)
	return nil
}
