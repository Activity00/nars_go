package storage

import (
	"nars_go/entity"
	"nars_go/language"
)

type Memory struct {
	Concepts *Bag[entity.Concept, language.Term]
}

func NewMemory() *Memory {
	return &Memory{
		Concepts: NewBag[entity.Concept, language.Term](),
	}
}

func (m *Memory) Cycle() {

}
