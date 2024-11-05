package primitives

import (
	"github.com/bearz-io/bzdev/lib/errors"
	"gopkg.in/yaml.v3"
)

type Task struct {
	Id          string    `json:"id"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Uses        string    `json:"uses"`
	With        *ExprsMap `json:"with"`
	Env         *ExprsMap `json:"env"`
	Cwd         *Expr     `json:"cwd"`
	Timeout     *Expr     `json:"timeout"`
	If          *Expr     `json:"if"`
	Needs       []string  `json:"needs"`
}

func (t *Task) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.MappingNode {
		return errors.NewYamlError(node, "task", "task must be a mapping node")
	}

	for i := 0; i < len(node.Content); i += 2 {

		keyNode := node.Content[i]
		valueNode := node.Content[i+1]

		switch keyNode.Value {
		case "id":
			if err := valueNode.Decode(&t.Id); err != nil {
				return err
			}
		case "name":
			if err := valueNode.Decode(&t.Name); err != nil {
				return err
			}
		case "description":
			if err := valueNode.Decode(&t.Description); err != nil {
				return err
			}
		case "uses":
			if err := valueNode.Decode(&t.Uses); err != nil {
				return err
			}
		case "with":
			t.With = &ExprsMap{}
			if err := valueNode.Decode(t.With); err != nil {
				return err
			}
		case "env":
			t.Env = &ExprsMap{}
			if err := valueNode.Decode(t.Env); err != nil {
				return err
			}
		case "cwd":
			t.Cwd = &Expr{}
			if err := valueNode.Decode(t.Cwd); err != nil {
				return err
			}
		case "timeout":
			t.Timeout = &Expr{}
			if err := valueNode.Decode(t.Timeout); err != nil {
				return err
			}
		case "if":
			t.If = &Expr{}
			if err := valueNode.Decode(t.If); err != nil {
				return err
			}
		case "needs":
			if err := valueNode.Decode(&t.Needs); err != nil {
				return err
			}
		}
	}

	return nil
}
