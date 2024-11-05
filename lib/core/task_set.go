package core

import "github.com/bearz-io/bzdev/schemas/primitives"

type TaskSet struct {
	kv    map[string]primitives.Task
	index []primitives.Task
}

func NewTaskSet() *TaskSet {
	return &TaskSet{
		kv:    make(map[string]primitives.Task),
		index: make([]primitives.Task, 0),
	}
}

func (ts *TaskSet) Len() int {
	return len(ts.index)
}

func (ts *TaskSet) Add(task primitives.Task) {
	if _, ok := ts.kv[task.Id]; ok {
		return
	}

	ts.kv[task.Id] = task
	ts.index = append(ts.index, task)
}

func (ts *TaskSet) Get(id string) (primitives.Task, bool) {
	task, ok := ts.kv[id]
	return task, ok
}

func (ts *TaskSet) GetAll() []primitives.Task {
	return ts.index
}

func (ts *TaskSet) Remove(id string) {
	if _, ok := ts.kv[id]; !ok {
		return
	}

	delete(ts.kv, id)
	for i, task := range ts.index {
		if task.Id == id {
			ts.index = append(ts.index[:i], ts.index[i+1:]...)
			break
		}
	}
}

// implement iterator
func (ts *TaskSet) Iterate() <-chan primitives.Task {
	ch := make(chan primitives.Task)
	go func() {
		for _, task := range ts.index {
			ch <- task
		}
		close(ch)
	}()
	return ch
}
