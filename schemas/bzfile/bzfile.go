package bzfile

import "github.com/bearz-io/bzdev/schemas/primitives"

type BzFile struct {
	Id    string                     `json:"id"`
	Name  *string                    `json:"name"`
	Tasks map[string]primitives.Task `json:"tasks"`
	Env   *primitives.ExprsMap       `json:"env"`
}
