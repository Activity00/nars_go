package storage

import "nars_go/entity"

type Bag[T, K any] struct {
	/** mapping from key to item */
	nameTable map[entity.Item[K]]T
}

func NewBag[T, K any]() *Bag[T, K] {
	return &Bag[T, K]{
		nameTable: make(map[entity.Item[K]]T, 0),
	}
}

func (b *Bag[T, K]) TakeOut() *T {
	if len(b.nameTable) == 0 {
		return nil
	}
	for _, v := range b.nameTable {
		return &v
	}

	return nil
}

func (b *Bag[T, K]) PutBack(key entity.Item[K], item T) {
	b.nameTable[key] = item
}
