package entity

import "fmt"

type BudgetValueOption func(*BudgetValue)

type BudgetValue struct {
	priority, durability, quality float64
}

func NewBudgetValue(option ...BudgetValueOption) *BudgetValue {
	value := new(BudgetValue)
	for _, o := range option {
		o(value)
	}

	return value
}

func (b *BudgetValue) String() string {
	return fmt.Sprintf("%f %f %f", b.priority, b.durability, b.quality)
}
