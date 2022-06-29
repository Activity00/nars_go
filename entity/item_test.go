package entity

import "testing"

func TestNewItem(t *testing.T) {
	item := NewItem(func(item *Item) {
		item.key = ""
		item.budget = NewBudgetValue()
	})
	item.setBudget()
}
