package core

import "context"

type EnvContext struct {
	Env     map[string]string
	Secrets map[string]string
	Vars    map[string]string
	Context context.Context
}

func NewEnvContext() *EnvContext {
	return &EnvContext{
		Env:     make(map[string]string),
		Secrets: make(map[string]string),
		Vars:    make(map[string]string),
		Context: context.Background(),
	}
}

type TaskContext struct {
	EnvContext *EnvContext
}
