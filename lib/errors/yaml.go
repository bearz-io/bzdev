package errors

import "gopkg.in/yaml.v3"

type YamlError struct {
	message     string
	node        *yaml.Node
	element     string
	description string
	code        string
}

func (e *YamlError) Error() string {
	return e.message
}

func (e *YamlError) Node() *yaml.Node {
	return e.node
}

func (e *YamlError) Element() string {
	return e.element
}

func NewYamlError(node *yaml.Node, element string, description string) error {
	return &YamlError{
		message:     "YAML error",
		node:        node,
		element:     element,
		description: description,
		code:        "YAML_ERROR",
	}
}
