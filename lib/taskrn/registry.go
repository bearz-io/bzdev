package taskrn

import "github.com/bearz-io/bzdev/schemas/primitives"

type TaskRegstry struct {
	data map[string]func(primitives.Task)
}
