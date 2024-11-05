package core

type TaskRegstry struct {
	data map[string]func(ctx *TaskContext) (Outputs, error)
}
