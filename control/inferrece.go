package control

import (
	"nars_go/entity"
	"nars_go/storage"
)

type InferenceControl struct{}

// Inference
// 1. get a concept from the memory
// 2. get a task from the concept
// 3. get a belief from the concept
// 4. derive new tasks from the selected task and belief
// 5. put the involved items back into the corresponding bags
// 6. put the new tasks into the corresponding bags
func Inference(mem *storage.Memory) {
	// 1. get a concept from the memory
	currConcept := mem.Concepts.TakeOut()
	if currConcept == nil {
		return
	}
	// 2. get a task from the concept
	task := currConcept.GetTask()

	// 3. get a belief from the concept
	belief := currConcept.GetBelief()
	// 4. derive new tasks from the selected task and belief
	tasks := DeriveNewTasks(task, belief)
	// 5. put the involved items back into the corresponding bags
	if currConcept != nil {
		mem.Concepts.PutBack(nil, *currConcept) // TODO
	}
	// 6. put the new tasks into the corresponding bags
	// mem.Concepts.PutBack(task)
	tasks = tasks
}

func DeriveNewTasks(task entity.Task, belief entity.Sentence) []string {
	return []string{}
}
