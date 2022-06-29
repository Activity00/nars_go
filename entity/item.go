package entity

import "fmt"

/**
 * An item is an object that can be put into a Bag,
 * to participate in the resource competition of the system.
 * <p>
 * It has a key and a budget. Cannot be cloned
 */

type ItemOption func(*Item)

type Itemer interface {
}

type Item struct {
	key    string
	budget *BudgetValue
}

func NewItem(option ...ItemOption) *Item {
	item := new(Item)
	for _, o := range option {
		o(item)
	}

	return item
}

func (i *Item) setBudget(budget *BudgetValue) {
	i.budget = budget
}

func (i *Item) GetKey(budget *BudgetValue) string {
	return i.key
}

func (i *Item) GetBudget() *BudgetValue {
	return i.budget
}

func (i *Item) String() string {
	return fmt.Sprintf("%s %s", i.budget.String(), i.key)
}
