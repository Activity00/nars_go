package storage

import "nars_go/entity"

type Bag[T entity.Itemer] struct {
	/** mapping from key to item */
	nameTable map[entity.Item]T
}

func NewBag[T any]() *Bag[T] {
	return &Bag[T]{
		nameTable: make(map[entity.Item]T, 0),
	}
}

func (b *Bag[T]) TakeOut() *T {
	if len(b.nameTable) == 0 {
		return nil
	}
	for _, v := range b.nameTable {
		return &v
	}

	return nil
}

func (b *Bag[T]) PutBack(key entity.Item, item T) {
	b.nameTable[key] = item
}
