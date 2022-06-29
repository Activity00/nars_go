package entity

type Concept struct {
}

func (c *Concept) GetTask() Task {
	return Task{}
}

func (c *Concept) GetBelief() Sentence {
	return Sentence{}
}
